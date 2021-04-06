# go-seleniumWD-godog-cucumber

**Correr jar file**

```
> java -jar resources/selenium-server-standalone-3.141.59.jar
```

**Reporte en formato .json**

```
> godog  --format=cucumber > log/report.json`
```

**Convertir reporte a formato html**

```
npm init

npm install cucumber-html-reporter --save-dev
```

creo archivo reporter.js

Creo archivo log/report.json con el comando

`$godog --format=cucumber > log/report.json`

Para ver el reporte

`node reporter.js`
