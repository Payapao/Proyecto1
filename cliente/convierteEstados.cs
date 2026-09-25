using System;

public class ConvierteEstados{

    public string toStringEstados(Estados e){
	string s = "";
	switch (e) {
	    case Estados.ACTIVE:
		s = "ACTIVE";
		break;
	    case Estados.AWAY:
		s = "AWAY";
		break;
	    case Estados.BUSY:
		s = "BUSY";
		break;
	    default:
		break;
	}
	return s;
    }
    
    public Estados toEstados(string s){
	Estados e = 0;
	switch (s) {
	    case "ACTIVE":
		e = Estados.ACTIVE;
		break;
	    case "AWAY":
		e = Estados.AWAY;
		break;
	    case "BUSY":
		e = Estados.BUSY;
		break;
	    default:
		break;
	}
	return e;
    }
    
}
