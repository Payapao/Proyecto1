using System;

public enum Estados{
    AWAY,
    ACTIVE,
    BUSY
}

public string toStringEstados(Estado e){
    string s = "";
    switch (e) {
	case ACTIVE:
	    s = "ACTIVE";
	    break;
	case AWAY:
	    s = "AWAY";
	    break;
	case BUSY:
	    s = "BUSY";
	    break;
	default:
    }
    return s;
}

public Estado toEstados(string s){
    Estado e = null;
    switch (s) {
	case "ACTIVE":
	    e = ACTIVE;
	    break;
	case "AWAY":
	    e = AWAY;
	    break;
	case "BUSY":
	    e = BUSY;
	    break;
	default:
    }
    return e;
}
