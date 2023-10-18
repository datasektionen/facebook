# Facebook
because it's *META*... you know?

### TODO
- bygg hela hemisdan ;-;
- Lägg till städschema

### DOCS (ish)

Sidan är just nu väldigt mycket i teststadiet men här är några pointers över struktur mm:

#### Client

- För att köra, använd ```bun run dev``` i mappen ```/client```
- Routing görs genom att skapa en ny fil med namnet av routen i mappen ```client/src/pages/```, exempelvis kommer ```<URL>/stadschema``` routa till default export från filen ```client/src/pages/stadschema.tsx```
- OBS: Den component som heter websockets är väldigt mycket för testande och inte kod som faktiskt ska finnas i sidan

#### Server

- För att köra, använd ````go run .``` i mappen ```/```
- Servern kräver en posgresql databas. Anslutningslänken är: ```postgres://facebook_user:facebook@localhost:5432/facebook``` sätt upp en databas runt dessa kriterier borde det lösa sig ;-;
- Självaste strukturen på databasen är väldigt mycket "in progress" och det schema som är definierat just nu är ingen bra lösning :)
- Alla actions som finns i servern är för testande, allstå finns inget riktigt backend API ännu 
