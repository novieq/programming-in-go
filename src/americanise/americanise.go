package americanise
import ( "bufio" //buffered io
	"fmt"
	"io"//ioreader and iowriter
	"io/ioutil" //high level file handling
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	inFileName, outFileName, err := filenamesFromCommandLine()
	if(err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	//Files is go are pointers to the type os.File, so we create two such variables initialized to standard input and output streams
	//which are both of type *os.File
	inFile, outFile := os.Stdin, os.Stdout

	//if the filename is empty both files have already been set to stdin and stdout
	//if the filenames are not empty then we create a new os file to read from or write to a file as appropriate
	if inFileName != nil {
		//os.open return os.file that can be used to read a file
		if inFile, err := os.Open(inFileName); err != nil {
			log.Fatal(err)
		}
		//the defer statement calls are executed when the function in which these calls are made actually return whether
		//normally or due to panic
		defer inFile.Close()

	}

	if outFileName != nil {
		//os.create returns os.file that can be used to read or write to a file, creating a new one if it doesnt exist and truncating
		// to 0 lines if it does exist
		if outFile, err := os.Create(outFileName); err!= nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}
	/*The americanise() function accepts an io.Reader and an io.Writer, not *os.Files, but this doesn’t matter since the os.File type supports the io.ReadWriter inter- face (which simply aggregates the io.Reader and io.Writer interfaces) and can therefore be used wherever an io.Reader or an io.Writer is required. This is an example of duck typing in action—the americanise() function’s parameters are interfaces, so the function will accept any values—no matter what their types—that satisfy the interfaces,*/
	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)

	}

}

//returns two strings and an error value
//return variables are set to 0 values when the function is entered and keep their zero values until explicitly assigned
func filenamesFromCommanLine() (inFileName, outFileName string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		/*The fmt.Errorf() function is like the fmt.Printf() function we saw earlier, except that it returns an error value containing a string using the given format string and arguments rather than writing a string to os.Stdout. (The errors.New() function is used to create an error given a literal string.)
*/
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt",
		filepath.Base(os.Args[0])) return "", "", err
	}
	if len(os.Args) > 1 {
		inFileName = os.Args[1] if len(os.Args) > 2 {
			outFileName = os.Args[2]
		}
	}
	if inFileName != "" && inFileName == outFileName {
		log.Fatal("won't overwrite the infile")
	}
	/*If a function or method has variable names as well as types listed for its return values, then a bare return is legal (i.e., a return statement that does not specify any variables). In such cases, the listed variables’ values are returned*/
	return inFileName, outFileName, nil
}