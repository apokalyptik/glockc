package glockc

import(
	"testing"
)

var client1 Client
var client2 Client
var err error

func TestLockcConnect(t *testing.T) {
	client1, err = New("127.0.0.1", 9999)
	if err != nil {
		t.Errorf( "client1 connection failed: %+v", err )
		t.FailNow()
	}
	client2, err = New("127.0.0.1", 9999)
	if err != nil {
		t.Errorf( "client2 connection failed: %+v", err )
		t.FailNow()
	}
}

func TestGet(t *testing.T) {
	var g1, g2 int
	g1, err = client1.Get("testlock", false)
	if g1 != 1 || err != nil {
		t.Errorf( "client1 Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g2, err = client2.Get("testlock", false)
	if g2 != 0 || err != nil {
		t.Errorf( "client2 Get failed: %d (%+v (%s))", g2, err, client2.DebugLast() )
	} else {
		t.Log( client2.DebugLast() )
	}
	g1, err = client1.Get("testlock", false)
	if g1 != 1 || err != nil {
		t.Errorf( "client1 Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
}

func TestInspect(t *testing.T) {
	var g1 int
	g1, err = client1.Inspect("testlock", false)
	if g1 != 1 || err != nil {
		t.Errorf( "client1 Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g1, err = client1.Inspect("neverlocked", false)
	if g1 != 0 || err != nil {
		t.Errorf( "client1 Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
}

func TestRelease(t *testing.T) {
	var g1 int
	g1, err = client1.Release("testlock", false)
	if g1 != 1 || err != nil {
		t.Errorf( "client1 Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g1, err = client1.Release("testlock", false)
	if g1 != 0 || err != nil {
		t.Errorf( "client1 Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g1, err = client1.Release("neverlocked", false)
	if g1 != 0 || err != nil {
		t.Errorf( "client1 Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
}

func TestName(t *testing.T) {
	g1, err := client1.Name("client1")
	if g1 != 1 || err != nil {
		t.Errorf( "client1 Name set failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
}

func TestSharedGet(t *testing.T) {
	var g1, g2 int
	g1, err = client1.Get("testlock", true)
	if g1 != 1 || err != nil {
		t.Errorf( "client1 shared Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g2, err = client2.Get("testlock", true)
	if g2 != 2 || err != nil {
		t.Errorf( "client2 shared Get failed: %d (%+v (%s))", g2, err, client2.DebugLast() )
	} else {
		t.Log( client2.DebugLast() )
	}
	g1, err = client1.Get("testlock", true)
	if g1 != 2 || err != nil {
		t.Errorf( "client1 shared Get failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
}

func TestSharedInspect(t *testing.T) {
	var g1 int
	g1, err = client1.Inspect("testlock", true)
	if g1 != 2 || err != nil {
		t.Errorf( "client1 shared Inspect failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g1, err = client1.Inspect("neverlocked", true)
	if g1 != 0 || err != nil {
		t.Errorf( "client1 shared Inspect failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
}

func TestSharedRelease(t *testing.T) {
	var g1 int
	g1, err = client1.Release("testlock", true)
	if g1 != 1 || err != nil {
		t.Errorf( "client1 shared Release failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g1, err = client1.Release("testlock", true)
	if g1 != 0 || err != nil {
		t.Errorf( "client1 shared Release failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g1, err = client1.Inspect("testlock", true)
	if g1 != 1 || err != nil {
		t.Errorf( "client1 shared Inspect failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
	g1, err = client1.Release("neverlocked", true)
	if g1 != 0 || err != nil {
		t.Errorf( "client1 shared Release failed: %d (%+v (%s))", g1, err, client1.DebugLast() )
	} else {
		t.Log( client1.DebugLast() )
	}
}
