package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// AdapterType represents the root element of the XML file.
/**
golang 语言写一个 解析 ice 61499 的 xml文件,xml 文件内容之前的回话已经给出来了。
1 将xml中的输入事件 转换成 两个c++ 函数 ，函数名称是分别是 "REQ","RSP"
2 "REQ" 函数 入参是事件 EventInputs REQ 对应的With Var 和 入参 EventOutputs Event Name="IND" 对应的With Var 指针。  并且 入参类型都转成小写字符。 返回值是bool类型 ，REQ 实现代码为空。
3 "RSP" 函数 入参是事件 EventInputs RSP 对应的With Var 和 入参 EventOutputs Event Name="CNF" 对应的With Var 指针。并且 入参类型都转成小写字符。返回值是bool类型 ，RSP 实现代码为空。。
4 并且包含一个main 对生成输入事件的调用，参数和函数入参类型保持一致。
5 输出C++文件。xml用读取文件路径的模式
6 golang 程序，代码有注释，模块清晰，易懂
*/
type AdapterType struct {
	GUID           string         `xml:"GUID,attr"`
	Name           string         `xml:"Name,attr"`
	Comment        string         `xml:"Comment,attr"`
	Namespace      string         `xml:"Namespace,attr"`
	Identification Identification `xml:"Identification"`
	VersionInfo    VersionInfo    `xml:"VersionInfo"`
	InterfaceList  InterfaceList  `xml:"InterfaceList"`
	Service        Service        `xml:"Service"`
}

// Identification contains standard information about the adapter.
type Identification struct {
	Standard string `xml:"Standard,attr"`
}

// VersionInfo contains versioning details of the adapter.
type VersionInfo struct {
	Organization string `xml:"Organization,attr"`
	Version      string `xml:"Version,attr"`
	Author       string `xml:"Author,attr"`
	Date         string `xml:"Date,attr"`
}

// InterfaceList contains input and output events and variable declarations.
type InterfaceList struct {
	EventInputs  []EventInput     `xml:"EventInputs>Event"`
	EventOutputs []EventOutput    `xml:"EventOutputs>Event"`
	InputVars    []VarDeclaration `xml:"InputVars>VarDeclaration"`
	OutputVars   []VarDeclaration `xml:"OutputVars>VarDeclaration"`
}

// EventInput represents an input event in the interface list.
type EventInput struct {
	Name     string    `xml:"Name,attr"`
	Comment  string    `xml:"Comment,attr"`
	WithVars []WithVar `xml:"With"`
}

// EventOutput represents an output event in the interface list.
type EventOutput struct {
	Name     string    `xml:"Name,attr"`
	Comment  string    `xml:"Comment,attr"`
	WithVars []WithVar `xml:"With"`
}

// WithVar represents a variable associated with an event.
type WithVar struct {
	Var string `xml:"Var,attr"`
}

// VarDeclaration declares a variable with its type and comment.
type VarDeclaration struct {
	Name    string `xml:"Name,attr"`
	Type    string `xml:"Type,attr"`
	Comment string `xml:"Comment,attr"`
}

// Service contains service details.
type Service struct {
	RightInterface   string            `xml:"RightInterface,attr"`
	LeftInterface    string            `xml:"LeftInterface,attr"`
	ServiceSequences []ServiceSequence `xml:"ServiceSequence"`
}

// ServiceSequence represents a sequence of service transactions.
type ServiceSequence struct {
	Name                string               `xml:"Name,attr"`
	ServiceTransactions []ServiceTransaction `xml:"ServiceTransaction"`
}

// ServiceTransaction represents a transaction within a service sequence.
type ServiceTransaction struct {
	InputPrimitive  InputPrimitive  `xml:"InputPrimitive"`
	OutputPrimitive OutputPrimitive `xml:"OutputPrimitive"`
}

// InputPrimitive represents an input primitive in a service transaction.
type InputPrimitive struct {
	Interface  string `xml:"Interface,attr"`
	Event      string `xml:"Event,attr"`
	Parameters string `xml:"Parameters,attr"`
}

// OutputPrimitive represents an output primitive in a service transaction.
type OutputPrimitive struct {
	Interface  string `xml:"Interface,attr"`
	Event      string `xml:"Event,attr"`
	Parameters string `xml:"Parameters,attr"`
}

// toLowerCase converts a string to lowercase.
func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// generateCPPCode generates C++ code based on the parsed XML data.
func generateCPPCode(adapter AdapterType) string {
	// Start building the C++ header code
	cppHeaderCode := fmt.Sprintf("#include <iostream>\n#include <string>\n\nusing namespace std;\n\nclass %s {\npublic:\n", adapter.Name)
	cppMainCode := fmt.Sprintf("int main() {\n")

	// Map to store variable names and their types
	varDeclarations := make(map[string]string)

	// Populate the map with input variables
	for _, varDecl := range adapter.InterfaceList.InputVars {
		varDeclarations[varDecl.Name] = toLowerCase(varDecl.Type)
	}
	// Populate the map with output variables
	for _, varDecl := range adapter.InterfaceList.OutputVars {
		varDeclarations[varDecl.Name] = toLowerCase(varDecl.Type)
	}

	// Build parameters for the REQ function
	reqFunctionParams := ""
	for _, event := range adapter.InterfaceList.EventInputs {
		if event.Name == "REQ" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				reqFunctionParams += fmt.Sprintf("%s %s, ", paramType, varName)
			}
		}
	}
	// Add IND output parameters to REQ function as pointers
	for _, event := range adapter.InterfaceList.EventOutputs {
		if event.Name == "IND" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				reqFunctionParams += fmt.Sprintf("%s* %s, ", paramType, varName)
			}
		}
	}

	// Build parameters for the RSP function
	rspFunctionParams := ""
	for _, event := range adapter.InterfaceList.EventInputs {
		if event.Name == "RSP" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				rspFunctionParams += fmt.Sprintf("%s %s, ", paramType, varName)
			}
		}
	}
	// Add CNF output parameters to RSP function as pointers
	for _, event := range adapter.InterfaceList.EventOutputs {
		if event.Name == "CNF" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				rspFunctionParams += fmt.Sprintf("%s* %s, ", paramType, varName)
			}
		}
	}

	// Remove trailing comma and space from function parameter lists
	if len(reqFunctionParams) > 0 {
		reqFunctionParams = reqFunctionParams[:len(reqFunctionParams)-2]
	}
	if len(rspFunctionParams) > 0 {
		rspFunctionParams = rspFunctionParams[:len(rspFunctionParams)-2]
	}

	// Define the REQ function prototype
	cppHeaderCode += fmt.Sprintf("    bool req(%s);\n", reqFunctionParams)
	// Define the RSP function prototype
	cppHeaderCode += fmt.Sprintf("    bool rsp(%s);\n", rspFunctionParams)
	// Close the class definition
	cppHeaderCode += "};\n\n"

	// Implement the empty REQ function
	cppHeaderCode += fmt.Sprintf("bool %s::req(%s) {\n", adapter.Name, reqFunctionParams)
	cppHeaderCode += "    // Implementation of req\n"
	cppHeaderCode += "    return true;\n}\n\n"

	// Implement the empty RSP function
	cppHeaderCode += fmt.Sprintf("bool %s::rsp(%s) {\n", adapter.Name, rspFunctionParams)
	cppHeaderCode += "    // Implementation of rsp\n"
	cppHeaderCode += "    return true;\n}\n\n"

	// Generate main function calls
	cppMainCode += fmt.Sprintf("    %s myComponent;\n", adapter.Name)

	// Initialize variables for REQ
	for _, event := range adapter.InterfaceList.EventInputs {
		if event.Name == "REQ" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				switch paramType {
				case "string":
					cppMainCode += fmt.Sprintf("    %s %s = \"Request Data\";\n", paramType, varName)
				case "int":
					cppMainCode += fmt.Sprintf("    %s %s = 123;\n", paramType, varName)
				default:
					cppMainCode += fmt.Sprintf("    %s %s = 0;\n", paramType, varName)
				}
			}
		}
	}

	// Initialize variables for IND output in REQ
	for _, event := range adapter.InterfaceList.EventOutputs {
		if event.Name == "IND" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				switch paramType {
				case "string":
					cppMainCode += fmt.Sprintf("    %s %s = \"Indication Data\";\n", paramType, varName)
				case "int":
					cppMainCode += fmt.Sprintf("    %s %s = 456;\n", paramType, varName)
				default:
					cppMainCode += fmt.Sprintf("    %s %s = 0;\n", paramType, varName)
				}
			}
		}
	}

	// Initialize variables for RSP
	for _, event := range adapter.InterfaceList.EventInputs {
		if event.Name == "RSP" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				switch paramType {
				case "string":
					cppMainCode += fmt.Sprintf("    %s %s = \"Response Data\";\n", paramType, varName)
				case "int":
					cppMainCode += fmt.Sprintf("    %s %s = 789;\n", paramType, varName)
				default:
					cppMainCode += fmt.Sprintf("    %s %s = 0;\n", paramType, varName)
				}
			}
		}
	}

	// Initialize variables for CNF output in RSP
	for _, event := range adapter.InterfaceList.EventOutputs {
		if event.Name == "CNF" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				paramType := varDeclarations[withVar.Var]
				switch paramType {
				case "string":
					cppMainCode += fmt.Sprintf("    %s %s = \"Confirmation Data\";\n", paramType, varName)
				case "int":
					cppMainCode += fmt.Sprintf("    %s %s = 101112;\n", paramType, varName)
				default:
					cppMainCode += fmt.Sprintf("    %s %s = 0;\n", paramType, varName)
				}
			}
		}
	}

	// Call REQ function
	cppMainCode += "\n    // Calling REQ function\n"
	cppMainCode += fmt.Sprintf("    myComponent.req(")
	firstParam := true
	for _, event := range adapter.InterfaceList.EventInputs {
		if event.Name == "REQ" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				if !firstParam {
					cppMainCode += ", "
				}
				cppMainCode += varName
				firstParam = false
			}
		}
	}
	for _, event := range adapter.InterfaceList.EventOutputs {
		if event.Name == "IND" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				if !firstParam {
					cppMainCode += ", "
				}
				cppMainCode += "&" + varName
				firstParam = false
			}
		}
	}
	cppMainCode += ");\n"

	// Call RSP function
	cppMainCode += "\n    // Calling RSP function\n"
	cppMainCode += fmt.Sprintf("    myComponent.rsp(")
	firstParam = true
	for _, event := range adapter.InterfaceList.EventInputs {
		if event.Name == "RSP" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				if !firstParam {
					cppMainCode += ", "
				}
				cppMainCode += varName
				firstParam = false
			}
		}
	}
	for _, event := range adapter.InterfaceList.EventOutputs {
		if event.Name == "CNF" {
			for _, withVar := range event.WithVars {
				varName := toLowerCase(withVar.Var)
				if !firstParam {
					cppMainCode += ", "
				}
				cppMainCode += "&" + varName
				firstParam = false
			}
		}
	}
	cppMainCode += ");\n"

	// End the main function
	cppMainCode += "    return 0;\n}\n"

	// Combine header and main code
	return cppHeaderCode + "\n" + cppMainCode
}

func main() {
	// Path to the XML file
	xmlFilePath := "ice_network.xml"

	// Read the XML file content
	xmlContent, err := ioutil.ReadFile(xmlFilePath)
	if err != nil {
		fmt.Printf("Error reading XML file: %v\n", err)
		return
	}

	// Unmarshal the XML content into an AdapterType struct
	var adapter AdapterType
	err = xml.Unmarshal(xmlContent, &adapter)
	if err != nil {
		fmt.Printf("Error unmarshalling XML: %v\n", err)
		return
	}

	// Generate C++ code from the parsed XML data
	cppHeaderCode := generateCPPCode(adapter)

	// Create or overwrite the generated C++ source file
	cppHeaderFile, err := os.Create("generated_component.cpp")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer cppHeaderFile.Close()

	// Write the generated C++ code to the file
	_, err = cppHeaderFile.WriteString(cppHeaderCode)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	// Print success message
	fmt.Println("Generated C++ source file successfully.")
}
