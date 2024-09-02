# UrlShortener

API funcional para acortar URL de manera sencilla hecha totalmente en Go utilizando Supabase como base de datos. Para ejecutarla habrá que configurar la variable de entorno DB_URL, que tendrá la cadena de conexión a una base de datos PostgreSQL (yo utilicé Supabase, tú puedes utilizar otra), y la variable PORT, que contendrá el puerto en el que se ejecutará la API. Para ello deberás crear un archivo .env en la raíz del proyecto y añadir lo siguiente:

```
DB_URL="postgresql://usuario:contraseña@host:puerto/base_de_datos"
PORT=":3000"
```

## Peticiones

### Crear
Para crear una ruta deberás realizar una petición a /url/create con el siguiente cuerpo:
``` json
{
  "custom_slug": "mi_custom_slug",
  "long_url": "https://mi-url.com/muy/muy/larga"
}
```

### Obtener todas las URLs
Deberás realizar una petición a /url y esta devolverá todas las URL creadas


### Obtener una URL específica
Al realizar una petición a /url/{slug} donde 'slug' debe ser el valor que indicamos mediante el campo 'custom_slug' al crear la entrada