package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

const bufsize int = 64 * 1024 * 1024

var buffer = make([]byte, bufsize, bufsize)
var buffer2 = make([]byte, bufsize, bufsize)

// Main
func main() {
	disk_dev := strings.TrimSpace(get_disk_dev())
	fmt.Printf("You entered device: %v\n", disk_dev)
	fmt.Printf("bufsize: %d\n", bufsize)
	prepare(0x55)
	fmt.Printf("buffer[%d] = 0x%02X, length = %d\n", 0, buffer[0:10],
		len(buffer))
	// firstPattern := string(buffer)
	// fmt.Println("First Pattern:", firstPattern[0:10])
	prepare(0xAA)
	fmt.Printf("buffer[%d] = 0x%02X, length = %d\n", 0, buffer[0:10],
		len(buffer))
	// secondPattern := string(buffer)
	// fmt.Println("Second Pattern:", secondPattern[0:10])
}

// Writes bytes to block device.
func write_file(file_name string) {
	var fd, num_bytes, total_bytes int
	var err error

	fmt.Printf("Writing file = %v. Block size = %v", file_name, bufsize)
	fd, err = syscall.Open(file_name, syscall.O_WRONLY, 0x00000600)
	check_error_object(err)

	total_bytes = 0
	for {
		num_bytes, err = syscall.Write(fd, buffer)

		total_bytes += num_bytes
	}
	err = syscall.Close(fd)
	check_error_object(err)
}

// Prepare buffer pattern
func prepare(pattern byte) {
	fmt.Printf("Preparing for phase 0x%02x\n", pattern)
	for i := 0; i < bufsize; i++ {
		buffer[i] = pattern
	}
}

// Get Disk Device
func get_disk_dev() string {
	reader := bufio.NewReader(os.Stdin)
	prompt := "Enter Disk Device (e.g., /dev/sda): "
	fmt.Print(prompt)

	input_dev, err_obj := reader.ReadString('\n')
	check_error_object(err_obj)
	if strings.Contains(input_dev, " ") {
		message := fmt.Sprintf("Device '%v' contains spaces", input_dev)
		panic(message)
	}
	// ToDo: Check if device begins with /dev/ and if device exists

	return input_dev
}

//Check error object
func check_error_object(err_obj error) {
	if err_obj != nil {
		panic(err_obj)
	}
}
