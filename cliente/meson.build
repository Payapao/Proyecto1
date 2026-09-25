project('cliente', 'cs', version: '1.3.2')

#Para la libreria de json
add_project_arguments('-r:System.Runtime.Serialization.dll', language: 'cs')

#Archivos del programa
archivos = files(
  'cliente.cs',
  'convierteEstados.cs',
  'convierteRespuestas.cs',
  'convierteTipo.cs',
  'estados.cs',
  'main.cs',
  'mensaje.cs',
  'respuestas.cs',
  'servidor.cs',
  'tipo.cs',
)
# Compila main 
executable('cliente', archivos)
