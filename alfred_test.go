package Alfred

import (
    // "fmt"
    "testing"
)

func TestBasics(t *testing.T) {
    tests := []struct {
        id       string
        expected string
    }{
        {id: "TestBasic", expected: "<items></items>"},
    }
    var ga *GoAlfred
    for _, test := range tests {
        ga = NewAlfred(test.id)
        res, err := ga.XML()
        sres := string(res)
        if err != nil {
            t.Fatalf("%s has faild with: %v", test.id, err)
        }
        if sres != test.expected {
            t.Errorf("Expected %v but received %v\n", test.expected, sres)
        }
    }
}

func TestAddItem(t *testing.T) {
    var ga *GoAlfred
    icon := NewIcon("pin.png", "icontype")
    ga = NewAlfred("TestAddItem")

    var tests = []struct {
        itemargs   []string
        make_valid bool
        expected   string
    }{
        {itemargs: []string{"uiduidadc", "TestBasic Title", "Adding stuff.", "yes", "yes", "file", "deleteme"},
            make_valid: false,
            expected: `<items>
  <item uid="uiduidadc" arg="deleteme" type="file" valid="yes" autocomplete="yes">
    <tittle>TestBasic Title</tittle>
    <subtitle>Adding stuff.</subtitle>
    <icon type="icontype">pin.png</icon>
  </item>
</items>`,
        },
        {itemargs: []string{"uiduidadc", "TestBasic Title", "Adding stuff.", "yes", "yes", "file", "deleteme"},
            make_valid: true,
            expected: `<items>
  <item uid="uiduidadc" arg="deleteme" type="file" valid="yes" autocomplete="yes">
    <tittle>TestBasic Title</tittle>
    <subtitle>Adding stuff.</subtitle>
    <icon type="icontype">pin.png</icon>
  </item>
  <item uid="uiduidadc" arg="" type="file" valid="no" autocomplete="yes">
    <tittle>TestBasic Title</tittle>
    <subtitle>Adding stuff.</subtitle>
    <icon type="icontype">pin.png</icon>
  </item>
</items>`,
        },
        {itemargs: []string{"", "", "Adding stuff.", "yes", "yes", "file", "deleteme"},
            make_valid: true,
            expected: `<items>
  <item uid="uiduidadc" arg="deleteme" type="file" valid="yes" autocomplete="yes">
    <tittle>TestBasic Title</tittle>
    <subtitle>Adding stuff.</subtitle>
    <icon type="icontype">pin.png</icon>
  </item>
  <item uid="uiduidadc" arg="" type="file" valid="no" autocomplete="yes">
    <tittle>TestBasic Title</tittle>
    <subtitle>Adding stuff.</subtitle>
    <icon type="icontype">pin.png</icon>
  </item>
  <item arg="" type="file" valid="no" autocomplete="yes">
    <tittle>No Result Were Found.</tittle>
    <subtitle>Adding stuff.</subtitle>
    <icon type="icontype">pin.png</icon>
  </item>
</items>`,
        },
    }
    for _, test := range tests {
        args := make([]string, 7)
        for i := 0; i < 7; i++ {
            args[i] = test.itemargs[i]
        }
        ga.AddItem(args[0], args[1], args[2], args[3], args[4], args[5],
            args[6], icon, test.make_valid)
        res, err := ga.XML()
        if err != nil {
            t.Fatalf("%s has faild with: %v", "TestAddItem", err)
        }
        if string(res) != test.expected {
            ferror(t, test.expected, string(res))
        }
    }
}

func TestMakeError(t *testing.T) {
    ga := NewAlfred("TestMakeError")
    ga.AddItem("uid", "title", "subtitle", "yes", "yes", "file", "arg",
        AlfredIcon{}, false)
    ga.MakeError(AlfredError("Testing Forcing an error result."))
    rec, err := ga.XML()
    if err != nil {
        t.Fatalf("%s has faild with: %v", "TestMakeError", err)
    }
    expected := `<items>
  <item arg="" valid="no" autocomplete="no">
    <tittle>Error in Generating Results.</tittle>
    <subtitle>Testing Forcing an error result.</subtitle>
    <icon>erroricon.png</icon>
  </item>
</items>`
    if string(rec) != expected {
        t.Errorf("Expected\n%v\nbut received ->\n%v\n", expected, string(rec))
    }
}

func ferror(t *testing.T, exp, rec interface{}) {
    t.Errorf("Expected\n%v\nbut received ->\n%v\n", exp.(string), rec.(string))
    t.Fail()
}
