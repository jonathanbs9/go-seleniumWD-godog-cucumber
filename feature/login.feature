    #language: en

    Funcionalidade: Login
    Para que yo pueda acceder a mis tareas
    Siendo un usuario
    Puedo autenticarme con mis datos previamente registrados

    Contexto: Formulario de login
        Dado que accedí a la página principal

    @success
    Scenario: Login de Usuario
        Cuando  hago el login con "jonathans@avalith.net" y "123jonathans"
        Entonces estoy autenticado con éxito

    Scenario: Contraseña incorrecta
        Cuando  hago el login con "jonathans@avalith.net" y "callefalsa123"
        Entonces debo ver el siguiente mensaje "El inicio de sesión no es válido."

    Scenario: Email inválido
        Cuando  hago el login con "rucurucu@avalith.net" y "123jonathans"
        Entonces debo ver el siguiente mensaje "El inicio de sesión no es válido."