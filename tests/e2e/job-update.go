package main

import (
	"bytes"
	"encoding/base64"
	"os"
	"strings"

	"github.com/dsnet/compress/bzip2"
)

func main() {
	job, err := os.ReadFile("./tests/e2e/job.yaml")
	if err != nil {
		panic(err)
	}
	jobString := string(job)
	// update kube config arg
	cc, err := os.ReadFile("./tests/e2e/kubeletconfig.json")
	if err != nil {
		panic(err)
	}
	cce, err := CompressAndEncode(cc)
	if err != nil {
		panic(err)
	}
	jobString = strings.ReplaceAll(jobString, "KUBELET_CONFIG", string(cce))

	// update kube config mapping arg
	cc, err = os.ReadFile("./tests/e2e/kubeletconfig-mapping.yaml")
	if err != nil {
		panic(err)
	}
	cce, err = CompressAndEncode(cc)
	if err != nil {
		panic(err)
	}
	jobString = strings.ReplaceAll(jobString, "KUBELET_MAPPING", string(cce))

	// update node config arg
	cc, err = os.ReadFile("./tests/e2e/nodeconfig.yaml")
	if err != nil {
		panic(err)
	}
	cce, err = CompressAndEncode(cc)
	if err != nil {
		panic(err)
	}
	jobString = strings.ReplaceAll(jobString, "NODE_CONFIG", string(cce))

	// update node commands arg
	cc, err = os.ReadFile("./tests/e2e/commands.yaml")
	if err != nil {
		panic(err)
	}
	cce, err = CompressAndEncode(cc)
	if err != nil {
		panic(err)
	}
	jobString = strings.ReplaceAll(jobString, "COMMANDS", string(cce))
	err = os.WriteFile("./tests/e2e/job.yaml", []byte(jobString), 0600)
	if err != nil {
		panic(err)
	}
}

func bzip2Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.DefaultCompression})
	if err != nil {
		return []byte{}, err
	}

	_, err = w.Write(data)
	if err != nil {
		return []byte{}, err
	}
	_ = w.Close()
	return buf.Bytes(), nil
}

func CompressAndEncode(data []byte) (string, error) {
	cm, err := bzip2Compress(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cm), nil
}
