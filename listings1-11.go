////////////////////////////////////////////////////////
//Listing 1: Definition und Verwendung eines Interface//
////////////////////////////////////////////////////////

type Stringer interface {
    String() string
}

func meinPrinter(s Stringer) {
        fmt.Println(s.String())
}

-------

/////////////////////////////////////
//Listing 2: time.Time als Stringer//
/////////////////////////////////////

func main() {  
	t := time.Now()
	meinPrinter(t)
}
// Output im Playground:
// 2009-11-10 23:00:00 +0000 UTC m=+0.000000001

-------

////////////////////////////////////////////////////////
//Listing 3: Definition des Interface bei der Funktion//
////////////////////////////////////////////////////////

type Loader interface {
	Load(UserKey) (User, error)
}

func CheckUser(uk UserKey, l Loader) error {
	u, err := l.Load(uk)
	if err != nil {
			return err
	}
	// ...
}

-------

////////////////////////////////////////////////////////
//Listing 4: Test-Implementierung des Loader-Interface//
////////////////////////////////////////////////////////

type LoadTester struct {
	users map[UserKey]User
}

func (l LoadTester) Load(uk UserKey) (User, error) {
	u, ok := users[uk]
	if !ok {
			return u, errors.New("User not found!")
	}
	return u, nil
}

-------

///////////////////////////////////////////////////////////////////////
//Listing 5: Verwenden eines Interface ohne darunterliegende Variable//
///////////////////////////////////////////////////////////////////////

var s Stringer // Wert ist nil
fmt.Println(s.String())
// panic: runtime error: invalid memory ...

-------

/////////////////////////////
//Listing 6: Auf nil prüfen//
/////////////////////////////

var s Stringer
if s != nil {
        fmt.Println(s.String())
}
// ...

-------

//////////////////////////////////
//Listing 7: Das leere Interface//
//////////////////////////////////

var a interface{}
var b int = 2
a = b
var c float64 = 3.1
a = c
var d string = "Hallo"
a = d

-------

/////////////////////////////
//Listing 8: Prüfen auf nil//
/////////////////////////////

err := openSomething(name)
if err != nil {
        // Fehlerbehandlung hier
}

-------

//////////////////////////
//Listing 9: Type-Switch//
//////////////////////////

func myFunc(s Stringer) {
	switch v := s.(type) {
	case nil:
			fmt.Println("nil Pointer")
	case *bytes.Buffer:
			fmt.Println("bytes.Buffer", v.Len())
	default:
			fmt.Println("Typ unbekannt")
}
}

-------

//////////////////////////////
//Listing 10: Type Assertion//
//////////////////////////////

func myFunc(r io.Reader) {
    buf, ok := r.(*bytes.Buffer)
    if ok {
        fmt.Println(buf.Bytes())
    }
}

-------

////////////////////////////////////
//Listing 11: Interface-Umwandlung//
////////////////////////////////////

func ReadAndClose(r io.Reader) ([]byte, error) {
	type closer interface {
			Close()
	}
	c, ok := r.(closer)
	if ok {
			defer c.Close()
	}
	return ioutil.ReadAll(r)
}