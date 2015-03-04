package pipeline

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_jenkins_multi_job_xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x4d\x73\xdb\x36\x10\xbd\xfb\x57\x70\x74\xf1\x49\x90\x9d\x36\x9d\x1c\x68\xe5\x43\x89\x67\xd2\x89\x15\x8d\x14\x8d\xcf\x10\xb9\x96\x60\x81\x00\x07\x00\x5b\xa9\x99\xfc\xf7\x2e\x40\x00\x04\x4d\xa5\xb5\xc6\x39\x99\xd8\x2f\x2c\xde\xbe\x7d\x72\xfe\xf6\x50\xf1\xec\x2f\x50\x9a\x49\x71\x73\x79\x4d\xae\x2e\x33\x10\x85\x2c\x99\xd8\xde\x5c\xae\xbf\xdd\x8e\xdf\x5c\xbe\x9d\x5e\xe4\x85\xac\x88\x61\x7b\xca\xc9\x23\x88\x3d\x13\x9a\xd4\xbc\xd9\xda\xbf\x55\xc3\x0d\x7b\x94\x1b\x72\x67\x3f\xfe\x94\x9b\x85\x92\x8f\x50\x98\xac\x0d\xb8\x19\xf9\x84\x71\x08\x1c\xb7\x8e\x77\xd7\xe4\xfa\xf5\x68\x7a\x91\x65\x39\x2d\x0c\xde\xae\x27\xee\x50\x82\x2e\x14\xab\xad\x65\x9a\x4f\xd2\x93\xf5\xee\x01\xea\x8f\x50\x83\x28\xb1\x4b\x06\x7a\xfa\x40\xb9\x86\x7c\x32\xb0\xdb\xe0\x5a\xc9\x1a\x94\xf1\x47\x34\x68\x20\x25\xa3\x1b\x99\xbc\x82\xd5\xc0\x99\x00\xb2\xf0\x1f\x8b\x36\xe7\x18\xdb\x2f\xd1\x8c\xf8\x1c\xc7\x21\x34\xf4\x7f\x45\xde\x90\x3f\x46\x6d\x65\xac\x6d\xa8\xde\xcf\x69\x05\xd3\x9a\x2a\xca\x39\xf0\x0c\x0e\x50\x34\xb6\xf3\x7c\x12\x9d\x21\x5a\x1b\xba\x05\x67\xf9\xfe\x3d\x23\xab\x70\xca\x7e\xfc\xc8\x27\x9d\xaf\xed\x7a\x72\x46\xdb\xee\xdd\x93\xfe\xc3\x73\x5d\x54\x59\xc1\xa9\xd6\x37\xa3\x5d\x53\x6a\x29\x08\x5a\xc8\xbc\xe1\x7c\x35\xbb\x1b\xb5\xb0\x17\x54\x2c\x25\xad\xa6\x46\x35\x08\x67\x38\xb9\x81\x30\x4d\x37\x1c\xca\x00\x75\x3c\x5b\xe7\x86\xcb\x62\xff\xa1\x61\xbc\xbc\xdf\x81\xf8\x28\xff\x16\xda\x28\xa0\x95\x33\x21\x87\x42\xd2\xff\xc6\x0d\x8b\xad\xeb\xe7\x94\x1a\x44\xd9\x42\x46\xb1\xed\x16\x39\xed\x9f\x26\x45\xd1\x28\x05\xc2\xb8\xa0\x50\xe7\xa9\xd9\x75\x60\xbf\x30\xd1\x23\x7f\x06\xeb\x3f\xb4\x99\x71\xc0\xf5\x8e\xea\x76\x88\xee\xeb\x0a\x87\x12\x2d\xbd\x18\xcc\xd5\xc1\x92\x65\xc8\x06\x45\xc5\x16\x90\x13\xcd\xc6\xba\x90\x11\xd1\xf9\x9c\x7e\x16\xa1\xe6\x4c\x8a\x07\xb6\xed\x2a\x63\x3a\xfa\x23\xe5\x1c\xd3\x82\x21\x0d\xb2\x98\x2c\x90\xc2\x95\x0e\x5c\xe8\x0c\x69\x1c\x1c\x6a\xa9\xa1\x44\x06\x05\x40\x13\x4b\x1a\xe8\xf9\x82\x3d\x3d\x61\x90\xb5\xa4\x81\xb8\x38\x1a\x27\xb8\x6c\x38\xe8\x05\x35\x3b\x5c\xff\x81\x29\x8d\xaf\xe8\x61\x09\x38\x6a\x64\x39\xc2\x9b\x9c\x7a\x6d\x0a\x7b\x95\xf5\x1c\x57\x46\x51\x03\xdb\x63\xec\xf7\x84\x6b\x98\x8a\x38\x96\xcc\xe9\x4f\x2f\xad\x33\xa7\x29\xb8\xa4\xca\xbc\xe7\xdc\x3e\xad\x45\x2f\xb5\xf4\x60\x8e\xf9\x8e\x89\x27\x6a\xed\x19\xe7\x6e\x9a\x5f\x05\x26\x2f\x41\xe3\x88\xbb\x5b\x6f\xdf\x7f\xfe\xb2\x5e\x7e\x42\xe5\xfb\xcf\xb0\xb4\xa0\x23\xf7\x57\xc1\x8f\x9f\x1f\x70\x46\xb3\x9d\xe5\x59\x14\xd0\xd3\xce\x8e\x7a\x93\x97\x70\x0f\x09\x87\xca\xdc\x51\xd9\x2f\x43\x4a\x7d\x0b\x88\x61\xa2\xa1\xb6\xed\xae\xff\xd5\x7a\x36\xfb\xb4\x5a\xdd\xae\xbf\x38\x98\x4e\x44\x5c\x3c\xbb\xbd\x13\xab\xea\xdf\xed\x37\x3e\xaf\x9b\x0d\x67\x7a\x17\x05\x00\xfb\x66\x0f\x19\x99\xc3\xc1\xdc\x51\xbc\x99\xa7\x1b\x99\xd3\x86\xd8\x5b\x0b\x54\x10\xd5\x54\xfa\xa8\x0d\x54\x9a\x78\x85\x6d\x6f\x27\xae\x7c\x54\x6b\xaf\x4b\xc4\x75\x10\xa4\xfb\x5b\x6b\x8c\x3f\x38\x2e\x65\xf0\x6b\x73\x4d\x7e\x27\xbf\x8d\x52\xb4\x10\xe1\x56\xe0\xda\x2d\x8b\x92\xea\x7f\x7f\xed\x62\x6b\xb7\xea\x83\xf6\x71\x03\x4f\x47\x7b\x2c\x7f\xed\xc3\x22\x94\x09\x05\x3a\x99\xb3\xcd\xf5\x50\xed\x5d\x83\x53\xb4\xca\x03\x06\x14\xfb\x07\xca\xde\x35\x4f\x71\xeb\x45\x8e\x7d\xe8\xbb\x57\xe4\xd5\xeb\x01\x6c\x09\xaf\xcf\xbd\xee\x84\xac\x0e\x8a\x3e\xbb\xf0\x5c\x96\xb0\x08\x8e\x6e\x98\x81\xd1\x83\xb2\xf6\x3f\x1a\x3b\x2d\xdd\x49\x78\xb4\x9c\x16\x17\xbf\x3f\x3f\xd3\x18\xdf\xc8\x3d\x33\xbb\xb9\xec\x3a\x09\xa2\xf0\x33\x77\x22\x0b\x2f\xc3\xef\xc9\x2b\xcf\x2f\x37\x24\x17\x42\xd2\xdb\xe2\x56\xf6\xee\x15\xad\xeb\x16\xe3\xb3\xb4\xc2\xaf\xc7\xf4\xe2\xdf\x00\x00\x00\xff\xff\x09\x99\x6b\x44\x29\x0b\x00\x00")

func templates_jenkins_multi_job_xml_bytes() ([]byte, error) {
	return bindata_read(
		_templates_jenkins_multi_job_xml,
		"templates/jenkins/multi-job.xml",
	)
}

func templates_jenkins_multi_job_xml() (*asset, error) {
	bytes, err := templates_jenkins_multi_job_xml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/jenkins/multi-job.xml", size: 2857, mode: os.FileMode(420), modTime: time.Unix(1423172303, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_jenkins_normal_job_xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x5f\x6f\xe3\xb8\x11\x7f\xf7\xa7\x10\xd2\x05\xd2\x1e\x6a\x79\x93\x6b\x71\x57\xc0\xf1\xed\xc6\xf1\xee\xb9\x48\xb2\x41\x6c\xef\x3e\x1e\x68\x69\x6c\xf3\x4c\x89\x02\x49\x65\xe3\xa6\xfb\xdd\x3b\xfc\x27\x51\x92\x9d\x75\xd0\x7b\x39\x20\xf0\x8b\x39\x33\x24\xe7\xff\xfc\xa8\xe1\x2f\x8f\x19\x8b\x1e\x40\x48\xca\xf3\x8b\xd3\xb3\xf8\xed\x69\x04\x79\xc2\x53\x9a\xaf\x2f\x4e\x17\xf3\x0f\xfd\x9f\x4f\x7f\x19\xf5\x86\x85\xe0\xbf\x43\xa2\x46\xbd\x28\x1a\x92\x44\xa1\xb0\x1c\x98\x45\x0a\x32\x11\xb4\xd0\x14\x4b\xd8\x02\x14\x57\x50\x40\x9e\xe2\x39\x14\xe4\x68\x45\x98\x84\xe1\xa0\x43\xd7\xc2\x78\x6c\x01\x42\xb9\x25\x12\x24\xc4\x29\x25\x4b\xce\xe2\xdf\x21\xdf\xd2\x5c\xc6\x05\x2d\x80\xd1\x1c\xe2\x3b\xf7\xe7\xce\xee\xd9\x45\x05\x2b\xd7\x34\xbf\x38\x49\x91\x8c\x16\xec\xfa\x5e\xb4\x6f\x39\xef\xde\xc6\x3f\xc7\x3f\x9d\xd8\x93\xf1\x6c\x45\xe4\xf6\x96\x64\x30\x7a\x7a\x8a\xe2\xb9\x5b\x44\xdf\xbe\x0d\x07\x15\xc7\x8b\x4a\x45\xd6\x50\xc9\xce\xfc\xca\x08\xd7\x3c\xab\xf2\xe0\x05\x3a\x1b\xa3\x07\x4d\xab\xf1\x06\xba\x8a\xe2\xa9\x9c\xe6\x54\x51\xc2\xfe\xcd\x97\x78\x8f\x16\x94\x49\x16\x25\x8c\x48\x79\x71\xb2\x29\x53\xc9\xf3\xd8\x1a\x26\xe3\x35\x55\xf1\x47\xaa\x66\xe3\x9b\x93\xca\x0d\x48\x7b\x77\x1e\x9f\xc7\xff\x72\x16\x0f\x13\x9e\xaf\xe8\xfa\xb3\x8d\xed\xe8\x7c\x38\x68\x12\xac\x50\x29\x41\xdc\x43\xc6\x15\x8c\x0d\x57\x56\x3e\xd8\x73\xe5\xa2\x25\xec\x65\xf5\x39\x82\x19\x5f\xa1\x56\x8b\xfb\x6b\xe3\x28\x4d\xf2\x87\x0d\x8e\x3e\x0d\xf7\xed\x57\x69\xb8\x14\x24\x4f\x36\xf0\xac\x86\x97\x46\x66\x56\x40\x12\xe8\x96\xeb\x60\xfd\x30\xd0\xea\xd5\x7c\xba\xa2\x20\x8c\x9e\x79\x18\xf9\x7d\x8a\xb6\x0f\x1d\x0e\x9a\xaa\x0c\x53\x2a\xc9\x92\xc1\xac\x5c\x66\x3c\x2d\x59\x9d\xf4\x5d\x86\xdd\x20\x20\x29\x31\x0a\x0f\x21\x47\x89\x12\x77\xec\xe3\xb8\x4b\xf8\x47\xc8\x41\x10\x55\xb3\xac\x7f\x4a\xa4\xe9\x82\xac\x2e\xfd\xae\xa0\x3d\x90\x94\x6a\xc3\xc5\x27\x31\xe6\x59\x46\x95\x02\xe1\x0f\xe8\x32\x5c\x42\x31\x20\xb9\x89\xf2\x58\xff\xfb\xc2\xc5\x56\x16\x24\xb1\x65\x61\x99\x56\xf0\x2b\xa6\xfd\xa7\x52\x55\x02\xfe\xe0\x0e\xdd\x8a\x17\xa2\xcc\xe1\xd2\xbb\xd4\xc9\x36\x89\xde\x6d\x3a\x2b\xee\x38\x63\x5e\x2a\xa0\x58\x11\xba\xce\xb9\x80\x5b\xae\xe8\x6a\x67\xf5\xf7\xa2\x7b\x38\x55\x09\xcc\x36\x84\x31\xfe\x75\xcc\x78\x0e\x2e\x10\x6d\xaa\xcb\xc1\x92\xb2\x74\xbc\xe1\x1c\x73\xf4\x99\xda\x2c\x15\x65\xf1\x15\xac\x48\xc9\xd4\x65\xb0\xe5\x64\xe0\xce\x41\xa1\x39\xe7\x6c\xe4\x64\x86\x03\x4f\x70\x6d\xb0\x0a\xdc\x6a\xed\xaf\x61\x54\xaa\x6a\xbf\x00\x86\xb1\x7c\x80\x39\x11\x6b\x50\x57\x54\xd4\x8c\x15\x08\xec\xb0\xe0\x09\xf0\x98\xb0\x32\x85\xf4\x1e\xd6\x55\xd7\x0e\xc8\xba\x08\x65\xa0\x95\xcd\x14\xdd\xde\x3a\xc4\x49\x46\x28\xf3\x54\xb9\xa5\xc5\x9c\xac\x9d\xb3\xfc\xca\xc5\x20\x6f\x5c\x39\x72\x1d\x4e\x87\x1d\xc7\x0a\x2a\x8b\x09\xa3\xb3\xa8\x4d\xc0\xd9\x60\x52\xa9\xbd\xdf\xdd\x98\x64\xa1\x5a\x7f\x78\x39\x7c\xd7\xe7\xf0\xa8\x20\x97\xf5\xe8\x1b\xa0\x46\xe6\x8f\x12\x74\xbd\x46\x37\x5a\x39\x67\xed\xa2\x90\x4a\x00\xc9\xb0\x9d\x4b\xdb\xcf\x51\xd2\x0f\x08\xbf\x23\xbe\x07\x3d\x7b\xc1\xe4\xc8\xdc\x12\xeb\x21\x84\x0d\x67\x50\xad\x4a\x77\xde\x9d\x9d\xc4\xc6\xab\xed\x4b\x30\x6b\xdb\x52\xd5\xf0\xdb\x08\x90\x1b\xce\xd2\x76\x5f\x9c\x2d\xc6\xe3\xc9\x6c\xd6\x6c\x83\xc8\xe4\x02\x31\x00\x61\xa3\xb7\xc3\x81\xff\x5b\x33\x13\xce\xb8\x18\x5d\x5e\x2f\x26\x7a\xaa\xe8\xff\x21\x2f\x2b\x18\x28\x6b\x93\xcb\x8f\x26\xad\xea\xb5\x2d\xa5\x86\x83\xe3\xfd\x53\xa5\x4b\x2f\x1c\x05\xd5\x3e\x9c\x8c\xfb\xdc\x39\xfa\x21\x72\x3f\x8c\x5e\x30\x24\x5c\xdb\xb8\xe3\x52\xd9\xd6\xf0\x2b\xe7\x5b\xd9\xec\x1c\x6d\x66\xaf\x31\x2f\x0e\xdc\x8c\x36\x06\xc9\x61\xdb\xc7\x17\x41\x8a\xa2\x4a\x97\xf6\x14\x23\x98\x62\xc6\xa5\xf1\x7b\xfc\x37\xd6\xff\x2e\x83\x5d\xd5\xb8\xaf\xe4\x10\xe6\xfc\x23\x7e\x5b\xc3\x1c\x43\xbc\x21\x85\x81\x28\x98\xb2\x22\x73\x31\xf2\xb4\xde\xde\x41\xf7\x9d\x7b\x8f\xc7\x67\x0e\x5f\x60\x8d\xa1\xe9\xcb\x52\x71\xf1\x62\xa4\xe6\xf0\xe8\x1c\x30\x6b\xb0\x6e\x47\x6f\x9e\x2e\x17\xd3\xeb\xab\xdf\x6e\x17\x37\x97\x93\xfb\x6f\xfd\x37\x4f\x1f\xa7\xf3\xdf\xee\x27\x9f\xa7\xb3\xe9\xa7\xdb\xbf\x33\xc8\xd7\x6a\x73\xf1\x13\xe6\x7f\x7b\x67\x5d\x3d\x29\x2e\xaf\xa8\x44\xf2\xce\x78\xc1\x75\xf9\x0e\xfd\x78\x50\xd7\x35\xd4\x46\xbc\x13\x64\x9d\xab\x98\x48\x7f\x42\x54\x67\x7a\x42\xa7\x35\xe0\x8c\x31\xc5\xa3\xc7\xef\x4a\x36\xb1\xd6\x2b\x0e\x7c\xc5\x81\xaf\x38\xf0\x15\x07\x46\x5d\x1c\x38\xd8\x03\xe4\x6a\x10\xf5\xe7\x1e\x8d\x07\xba\xbe\x45\x28\xf5\x23\x7f\xc6\xc8\x03\x5c\x93\x25\x30\x37\x0c\x30\xa2\x98\x8f\x90\xde\xf2\xd4\x7d\x6b\x08\x25\xb0\x04\x43\x7e\xaf\x09\x7b\x86\x09\xc9\xef\x39\xc9\x7c\x62\xfb\x65\xaf\xee\x44\x69\xab\x01\xa5\xd6\xd5\x8c\x27\x5b\x6b\xc1\x06\xf2\x2b\xfe\x35\xb7\xa0\xd1\x90\x10\x90\xfb\x4d\xdf\x95\xeb\x1e\xe6\x01\xe9\xf3\x47\x75\xa4\x7a\x76\xb0\x61\xcb\xc3\x9c\xb5\x95\x52\x59\xd5\x22\xf7\x9e\x41\xd8\x2f\xc0\xd7\x21\xba\xfe\xbf\xb0\x75\x17\x59\x3f\x83\xab\x9f\x41\xd5\x87\x31\xf5\xb1\x88\xba\x85\xa7\x8f\x47\xd3\x8d\xa4\x32\x79\x1c\x3e\x66\xb0\xf5\xae\x21\x8a\xdf\x0b\x6c\x98\x24\x51\x57\x50\xb4\x51\xb7\xaf\x93\x84\x17\x3b\xe2\xc4\xe2\x31\x2e\xfc\x9e\xaa\x32\x43\x89\x77\x67\xf1\x8f\xe7\xf1\x59\x5d\x9d\xfe\x1b\xa3\x76\xbe\xf3\xb1\x2e\x4a\x83\xe5\x4c\x00\x82\x8f\x90\x66\xc3\x8a\x32\x3d\x94\xb4\x7c\x75\x93\x96\x73\xf4\xea\xe9\x63\x7a\x63\xfd\x92\x72\x2d\x4f\xd6\x14\x09\x0c\xcf\xe5\x87\x3a\x79\xc3\x2e\xe7\x37\x48\x8d\x17\x67\x6e\xe7\x49\x00\x33\x30\x71\xd9\x92\x24\xdb\x39\xbf\x26\x52\xcd\xca\x24\x01\x29\x57\x25\x73\x91\x3b\xc8\x0e\x40\x94\xcb\xb4\x0f\xc6\x90\x99\xd2\x33\x7c\xbd\x1b\x61\x8f\xfe\xc8\xf8\x92\xb0\x19\x28\x85\x65\x53\xa7\x64\x4b\xb0\x02\x31\xde\xb0\x8a\x92\x72\x1c\x7c\x1f\x70\x2f\x88\x42\xd0\x5c\x79\xbf\x05\xa0\xe1\x90\xc0\xfe\xde\x78\x30\xe6\xcf\x3e\xd3\x88\xdc\xe2\x4b\x69\x03\x2c\x4c\xff\x2c\x23\x39\x66\xee\x5f\x22\x8f\xb4\xa3\x04\x6d\xd3\x61\xc9\xc8\x9a\x26\x3d\x78\x2c\xb8\x50\xd1\xdd\xf4\x6e\x72\x3d\xbd\x9d\xb8\xd7\xc0\xc5\x9b\xbf\x42\xb2\xe1\xd1\xc9\x9b\xa7\x8a\xf3\x79\x72\xaf\x1f\x06\xdf\x4e\xa2\xff\x46\x49\xa9\xa2\xfe\xea\x2c\xea\xa7\xa7\xfd\xd3\xbf\xf9\x43\xf4\xfb\x61\xf6\xeb\xfb\xe3\x36\x9f\xbb\xcd\xfb\xbf\xd4\x06\xe8\x1e\x47\x64\x84\x15\x08\xb8\xa9\xbf\x21\x22\x8d\xec\x43\x05\x2f\xaa\x3f\x6f\xf4\x0c\x6e\xb2\xc6\x7a\xbf\x68\xbf\x56\xf6\x37\xdc\xdc\xf2\x94\x1b\x33\xfe\x39\x59\x94\x4b\x44\x03\x9b\xce\xc0\xb4\xbb\x6e\x70\x4c\x07\x53\x31\x33\x4b\xac\xba\xb3\x60\x22\x22\xcc\xa4\x05\xc5\xe6\x1a\x54\x43\x8a\x0f\x1a\x0b\x8f\x26\xfa\xb1\xb6\xc0\xae\xaf\x67\x47\xd8\x7d\x9e\x15\xa9\xab\x2a\x4f\xe7\x7c\x9a\xa7\xf4\x81\xa6\x25\xa6\x97\x07\x09\x1d\xfa\x3e\x9b\xad\xf6\x8d\x4f\x2a\x41\x91\xef\x31\xd7\x73\xdf\x8b\x64\xa3\x5f\x99\x95\x1e\xa4\x4a\xe1\x4e\xa7\x20\xcd\xec\xd6\x0d\x3a\x67\xbb\xe9\xaa\x53\xb4\x1d\xb2\x97\xd7\xaf\x4c\xa9\x3e\x21\xdb\x49\x06\x84\x4a\x01\x8d\x1e\x27\x59\xa1\x76\x4e\xb7\x0a\x60\x77\x18\xfb\x3c\xb1\xdf\xb0\x66\x65\x39\x0f\xdd\xc2\xa3\xba\x21\x79\x69\x72\xb3\xfe\xec\x44\x4a\xac\xd3\x2c\x4e\x30\xce\xa2\xcc\xe4\x4e\x2a\xc8\x64\xdc\xa8\xe4\xd8\x24\x56\xf5\xc6\x75\x23\x23\x36\x11\xf5\x0f\x5e\xd7\xfa\xaa\x8c\x32\x5b\x3a\x6f\xf9\x33\xc4\x5c\x3f\x86\x98\xcb\xbc\xdd\xc2\xf4\xf2\x38\x22\x68\xf2\x36\x36\x1d\xf5\x75\xa6\xed\x95\x76\x7e\xfa\x63\x0d\x3b\xe0\x58\x37\x00\xb5\x72\x0d\xaf\xb6\x3a\x61\x41\x04\xaa\x86\x6d\x98\xfe\x07\xd2\xc6\x35\x6d\xbf\x35\x24\xfb\x4e\x54\xbf\xf2\xff\xd9\x71\x5b\x30\x14\x5e\x7a\x5d\xfb\x31\xbf\xef\xd0\xa3\x0f\xd6\xf8\xf3\xce\x33\xea\x60\xfa\xd6\xd5\x39\xd6\x0f\x73\x1b\xd7\x70\x7a\xcb\xb6\x46\x29\xd5\x8f\xd1\x1a\x2f\xd5\xa4\x50\xd0\x29\xf2\x85\xaa\xcd\x2d\xaf\x35\xf1\x95\x74\x88\x5d\xbb\xaf\x3d\xb8\x5e\xe8\xbf\x96\x95\x2f\x3f\xae\x9b\x5c\xe8\x92\xa0\x7f\x07\xf0\xe6\x7f\x01\x00\x00\xff\xff\xa3\x99\x23\x86\x95\x1d\x00\x00")

func templates_jenkins_normal_job_xml_bytes() ([]byte, error) {
	return bindata_read(
		_templates_jenkins_normal_job_xml,
		"templates/jenkins/normal-job.xml",
	)
}

func templates_jenkins_normal_job_xml() (*asset, error) {
	bytes, err := templates_jenkins_normal_job_xml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/jenkins/normal-job.xml", size: 7573, mode: os.FileMode(420), modTime: time.Unix(1425474447, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_jenkins_pipeline_xml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\xcb\x6e\xdb\x30\x10\xbc\xfb\x2b\x08\xa1\x57\x53\x4e\x0f\x45\x0e\x34\xd3\xc2\x6d\x80\x16\x7d\xa4\x68\xda\x6b\xc1\x48\x6b\x9a\x2d\xb5\x24\xf8\xb0\x1d\x04\xf9\xf7\xae\x1e\x4c\xa2\xf4\x01\x04\x3d\x49\x9c\x99\x9d\x9d\xd5\x52\xe2\xec\xd8\x59\xb6\x87\x10\x8d\xc3\x75\x75\xc2\x57\x15\x03\x6c\x5c\x6b\x50\xaf\xab\xaf\x97\xe7\xcb\xd3\xea\x4c\x2e\x44\x04\xde\x1a\x75\xe5\x2c\xff\x01\xf8\xd3\x60\xe4\xde\x78\xb0\x06\x81\xbf\xa6\x07\x19\x5c\x5f\x4c\xc0\x37\x03\x07\xe6\x6d\xd6\x86\x0c\xdb\x89\x5c\x16\xf9\x72\x64\x5e\xae\xf8\x29\x7f\x51\xc9\x05\x63\x02\x55\x07\xf2\xe6\x86\xf1\x8f\xf4\xc2\x6e\x6f\x45\x3d\x20\x3d\xb5\x35\x36\x41\x78\x73\x84\x26\x27\x17\xa2\xdc\x2a\x1b\x41\xd4\x8f\xe1\x7b\xe9\xe7\x0c\x19\xe6\xb2\x11\xea\x25\x3e\x38\x0f\x21\x19\x88\xac\xb1\x2a\xc6\x75\xb5\xcb\x6d\x74\xc8\x3b\x47\x39\x79\x1f\xfc\xd9\xc5\xa8\xb9\x7e\x6f\x62\xaa\xea\xa1\xac\x71\x9d\x77\x08\x98\xbe\x78\x68\x86\x66\x04\x3e\xf1\x83\x7c\x5f\x6e\x1e\xba\x8c\x26\xff\x9c\x7d\xa0\xb7\x26\xc4\xf4\xce\x5d\xcd\x25\x77\xe8\x18\xa5\xfe\xff\x2c\xa2\xfe\x7d\x48\x81\xee\xd3\xb6\xd4\x45\x79\xb2\xa2\x6c\x33\xa4\xd7\xc4\x9d\x3b\xbc\xd2\x3a\x80\x56\x09\xda\xc2\x95\x05\xfc\x85\x2d\xe6\x1b\x67\x73\x87\x64\x3d\x3a\x97\xe3\xe0\xeb\x68\x4f\xa8\x25\x52\x26\xf2\x99\x4e\x77\x1d\xf7\x2a\xa9\xfb\xeb\xf0\x10\xea\x25\xd9\xb7\xd4\xef\x2d\xd2\xf2\xf7\xca\xca\xe7\xa2\x7e\x84\x14\x9f\xcd\x4e\xa1\x86\x99\x4f\x81\x7a\x89\xb2\xd6\x1d\x3e\x28\xcc\xca\x5e\x06\xa3\x35\xfd\x25\x32\x85\x4c\xca\x3f\x31\x7d\x05\x4d\x0a\x47\x7f\x3e\xed\x27\xd2\xfd\x79\xf2\x76\xe4\xe2\x57\x00\x00\x00\xff\xff\x5f\x40\xf2\x3f\x94\x03\x00\x00")

func templates_jenkins_pipeline_xml_bytes() ([]byte, error) {
	return bindata_read(
		_templates_jenkins_pipeline_xml,
		"templates/jenkins/pipeline.xml",
	)
}

func templates_jenkins_pipeline_xml() (*asset, error) {
	bytes, err := templates_jenkins_pipeline_xml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/jenkins/pipeline.xml", size: 916, mode: os.FileMode(420), modTime: time.Unix(1423648470, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/jenkins/multi-job.xml":  templates_jenkins_multi_job_xml,
	"templates/jenkins/normal-job.xml": templates_jenkins_normal_job_xml,
	"templates/jenkins/pipeline.xml":   templates_jenkins_pipeline_xml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"jenkins": &_bintree_t{nil, map[string]*_bintree_t{
			"multi-job.xml":  &_bintree_t{templates_jenkins_multi_job_xml, map[string]*_bintree_t{}},
			"normal-job.xml": &_bintree_t{templates_jenkins_normal_job_xml, map[string]*_bintree_t{}},
			"pipeline.xml":   &_bintree_t{templates_jenkins_pipeline_xml, map[string]*_bintree_t{}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
