using System;
using System.Text.Json;
using System.Text.Json.Serialization;

public class Mensaje {
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public Tipo Type {get; set;}
    
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]    
    public string Roomname {get; set;}
    
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public string Username {get; set;}
    
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public string[] Usernames {get; set;}

    [JsonConverter(typeof(JsonStringEnumConverter))]
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public Estados Status {get; set;}

    [JsonConverter(typeof(JsonStringEnumConverter))]
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public Tipo Operation {get; set;}

    [JsonConverter(typeof(JsonStringEnumConverter))]
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public Respuesta Result {get; set;}

    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public string Text {get; set;}

    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public Dictionary<string,Estados> Users {get; set;}

    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public string Extra {get; set;}

    
    public FabricaMensaje(Type t){
	this.Type = t;
    }

    public roomname(string r)(Mensaje m){
	this.Roomname = r;
	return this;
    }

    public username(string u)(Mensaje m){
	this.Username = u;
	return this;
    }

    public usernames(string[] u)(Mensaje m){
	this.Usernames = u;
	return this;
    }

    public status(Estado e)(Mensaje m){
	this.Status = e;
	return this;
    }

    public operation(Tipo o)(Mensaje m){
	this.Operation = o;
	return this;
    }

    public result(Respuesta r)(Mensaje m){
	this.Result = r;
	return this;
    }

    public text(string t)(Mensaje m){
	this.Text = t;
	return this;
    }

    public users(string[] u)(Mensaje m){
	this.Users = u;
	return this;
    }

    public extra(string e)(Mensaje m){
	this.Extra = e;
	return this;
    }
}
    
