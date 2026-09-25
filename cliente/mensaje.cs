using System;
using System.Collections.Generic;
using System.Runtime.Serialization;

[DataContract]
public class Mensaje {

    //Para poder definir como se van a imprimir las enunmeraciones
    public static ConvierteTipo ct = new ConvierteTipo();
    public static ConvierteRespuestas cr = new ConvierteRespuestas();
    public static ConvierteEstados ce = new ConvierteEstados();

    
    public Tipo Type {get; set;}
    
    [DataMember(Name = "type", EmitDefaultValue = false)]
    public string Tipo {
	get {return ct.toStringTipos(this.Type);}
	set {this.Type = ct.toTipo(value);}
    }
    
    [DataMember(EmitDefaultValue = false)]  
    public string Roomname {get; set;}
    
    [DataMember(Name = "username", EmitDefaultValue = false)]
    public string Username {get; set;}
    
    [DataMember(EmitDefaultValue = false)]
    public string[] Usernames {get; set;}

    public Estados Status {get; set;}
    
    [DataMember(Name = "status", EmitDefaultValue = false)]
    public string Estado{
	get {return ce.toStringEstados(this.Status);}
	set {this.Status = ce.toEstados(value);}
    }
    
    public Tipo Operation {get; set;}
    
    [DataMember(Name = "operation", EmitDefaultValue = false)]
    public string Operacion {
	get {return ct.toStringTipos(this.Operation);}
	set {this.Operation = ct.toTipo(value);}
    }
    
    public Respuesta Result {get; set;}
    
    [DataMember(Name = "result", EmitDefaultValue = false)]
    public string Respuesta {
	get {return cr.toStringRespuestas(this.Result);}
	set {this.Result = cr.toRespuesta(value);}
    }
    
    [DataMember(EmitDefaultValue = false)]
    public string Text {get; set;}
    
    public Dictionary<string,Estados> Users {get; set;}

    [DataMember(EmitDefaultValue = false)]
    private Dictionary<string, string> Usuarios {
        get{
            if (Users == null)
		return null;
            var d = new Dictionary<string, string>();
            foreach (var usuario in d) 
		//Para que los estados sean strings
                d[usuario.Key] = usuario.Value.ToString(); 
            return d;
        }
        set{
            if (value == null) {
                Users = null;
                return;
            }
            Users = new Dictionary<string, Estados>();
            foreach (var usuario in value) 
                if (Enum.TryParse<Estados>(usuario.Value, true, out Estados estado)) 
                    Users[usuario.Key] = estado;

	}
    }
    
    [DataMember(EmitDefaultValue = false)]
    public string Extra {get; set;}

    
    public static Mensaje FabricaMensaje(Tipo t){
	Mensaje m = new Mensaje();
	m.Type = t;
	return m;
    }

    public Mensaje roomname(string r){
	this.Roomname = r;
	return this;
    }

    public Mensaje username(string u){
	this.Username = u;
	return this;
    }

    public Mensaje usernames(string[] u){
	this.Usernames = u;
	return this;
    }

    public Mensaje status(Estados e){
	this.Status = e;
	return this;
    }

    public Mensaje operation(Tipo o){
	this.Operation = o;
	return this;
    }

    public Mensaje result(Respuesta r){
	this.Result = r;
	return this;
    }

    public Mensaje text(string t){
	this.Text = t;
	return this;
    }

    public Mensaje users(Dictionary<string, Estados> u){
	this.Users = u;
	return this;
    }

    public Mensaje extra(string e){
	this.Extra = e;
	return this;
    }
}
    
