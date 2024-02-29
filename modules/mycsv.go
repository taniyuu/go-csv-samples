package modules

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
	"golang.org/x/text/transform"
)

// ファイルを作成してCSV書き込みを実施(encoding/csv")
func EncodeStandard(fileName string, transformer transform.Transformer, data [][]string) error {
	f, err := os.Create(fileName)
	defer func() {
		if err := f.Close(); err != nil {
			log.Default().Print(err)
		}
	}()
	if err != nil {
		return err
	}
	cw := csv.NewWriter(transform.NewWriter(f, transformer))
	return cw.WriteAll(data)
}

// ファイルを作成してCSV書き込みを実施(encoding/csv")
func EncodeCSVUtil(fileName string, transformer transform.Transformer, data interface{}) error {
	f, err := os.Create(fileName)
	defer func() {
		if err := f.Close(); err != nil {
			log.Default().Print(err)
		}
	}()
	if err != nil {
		return err
	}
	cw := csv.NewWriter(transform.NewWriter(f, transformer))
	encoder := csvutil.NewEncoder(cw)
	if err := encoder.Encode(data); err != nil {
		return err
	}
	cw.Flush()
	return cw.Error()
}

// ファイルを作成してCSV書き込みを実施(encoding/csv")
func EncodeDoubleQuote(fileName string, transformer transform.Transformer, data interface{}) error {
	// 一旦全てエンコードする
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := csvutil.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	reader := csv.NewReader(&buf)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error:", err)
	}

	//rewrite to add "", escape \"
	bytes := make([]byte, 0, len(buf.Bytes())*2)

	//If you update writer by SetCSVWriter, please change the crlf which you use
	for _, line := range lines {
		fmt.Println(line)
		for i, part := range line {
			if i != 0 {
				bytes = append(bytes, byte(delimiter))
			}
			bytes = append(bytes, []byte(escape(part))...)
		}
		bytes = append(bytes, byte('\r'))
		fmt.Println(string(bytes))
		fmt.Println("-------")
	}

	// f, err := os.Create(fileName)
	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 		log.Default().Print(err)
	// 	}
	// }()
	// if err != nil {
	// 	return err
	// }
	cw := csv.NewWriter(transform.NewWriter(f, transformer))
	encoder := csvutil.NewEncoder(cw)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("error:", err)
	}
	cw.Flush()
	return cw.Error()
}

func escape(part string) string {
	//"XXX" => XXX
	escapeStr := strings.Replace(part, "\"", "\"\"", -1)
	return "\"" + escapeStr + "\""
}
