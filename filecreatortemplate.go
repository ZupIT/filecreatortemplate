/*
 * Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package filemanagertemplate

import (
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/gobuffalo/packr/v2"
)

func CreatFile(path string, fileName string, extension string, templateLocation string, templateName string, replace map[string]string, writer io.Writer) {
   
    box :=  packr.New("fileBox", templateLocation)
    data, err := box.FindString(templateName)
    var templateContent = string(data)

    if err != nil{
        panic("Error to read the template")
    }
    
    var fileCompleteName = path+"/"+fileName+"." + extension

    if fileExists(fileCompleteName){
        panic("File already exist")
    }

     destination, err := os.Create(fileCompleteName)
     if err != nil {
        panic(err)
        return
     }
     defer destination.Close()
     var content = createFileUsingTemplate(templateContent, replace)
     fmt.Fprintf(destination,content)
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func createFileUsingTemplate(templateContent string, replace map[string]string) string{
    var result = templateContent
    for key, value := range replace {
        result = strings.ReplaceAll(result, "{"+key+"}", value)
    }
    return result
}
