package filestore

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 存储分为两部分：（1）原始文件（2）文件信息
// 文件名全部以原始文件的MD5命名，比如文件A的MD5是xxx，
// 那么保存的原始文件为xxx.bin，其文件信息为xxx.info

type FileInfo struct {
	ID     string // 文件的ID，这里全部使用原始文件的MD5
	Type   string // 文件类型
	Size   int64  // 文件大小
	Offset int64  // 文件偏移，当Offset等于Size时，表明该文件是完整的
}

type FileStore struct {
	Path string // 文件保存路径
}

func NewFileStore(path string) *FileStore {
	return &FileStore{
		Path: strings.TrimRight(path, "/"),
	}
}

// 新建一个文件上传任务，传入id（即文件的MD5）
// 返回期望的接收的offset、文件是否已经完整、错误信息
// 若offset为0，则表示服务器不存在该文件，若offset不为0，则表示服务器存在该文件的部分或全部内容
func (fs *FileStore) NewUpload(info *FileInfo) (int64, bool, error) {
	data, err := ioutil.ReadFile(fs.infoPath(info.ID))
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.OpenFile(fs.binPath(info.ID), os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				return 0, false, err
			}
			file.Close()
			return 0, false, fs.WriteInfo(info.ID, info)

		} else {
			return 0, false, err
		}
	}

	existInfo := &FileInfo{}
	if err = json.Unmarshal(data, existInfo); err != nil {
		return 0, false, err
	}

	// 文件已完整存在
	if existInfo.Offset >= existInfo.Size {
		return existInfo.Offset, true, nil
	}
	return existInfo.Offset, false, nil
}

func (fs *FileStore) WriteInfo(id string, info *FileInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fs.infoPath(id), data, os.ModePerm)
}

func (fs *FileStore) GetInfo(id string) (*FileInfo, error) {
	data, err := ioutil.ReadFile(fs.infoPath(id))
	if err != nil {
		return nil, err
	}

	info := &FileInfo{}
	if json.Unmarshal(data, info) != nil {
		return nil, err
	}
	return info, nil
}

// 向文件（id）的偏移地址offset处写入数据
// 返回写入的字节数、文件是否已经完整、错误信息
func (fs *FileStore) WriteChunk(id string, offset int64, src io.Reader) (int64, bool, error) {
	info, err := fs.GetInfo(id)
	if err != nil {
		return 0, false, err
	}

	// 文件已完整存在
	if info.Offset >= info.Size {
		return 0, true, nil
	}

	// 若offset跟期望接收的offset不一样，则直接返回错误
	if offset != info.Offset {
		return 0, false, errors.New("offset is unexpected")
	}

	file, err := os.OpenFile(fs.binPath(id), os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return 0, false, err
	}
	defer file.Close()

	n, err := io.Copy(file, src)
	if err != nil {
		return n, false, err
	}

	if n > 0 {
		info.Offset = offset + n
		err = fs.WriteInfo(id, info)
	}

	if info.Offset >= info.Size {
		return n, true, err
	}
	return n, false, err
}

func (fs *FileStore) binPath(id string) string {
	return fs.Path + "/" + id + ".bin"
}

func (fs *FileStore) infoPath(id string) string {
	return fs.Path + "/" + id + ".info"
}
