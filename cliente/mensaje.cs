using System;
using System.Text.Json;
using System.Text.Json.Serialization;

public class Mensaje {
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    [JsonIgnoreCondition(JsonIgnoreConditionWhenWritingNull)]
    public Tipo Type {get; set;}
    
    
    public FabricaMensaje(Type t){
	this.Type = t;
    }
}
    
