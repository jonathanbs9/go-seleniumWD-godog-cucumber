Feature: Login
    Para que yo pueda acceder a mis tareas
    Siendo un usuario
    Puedo autenticarme con mis datos previamente registrados
    
    @success
    Scenario: Login de Usuario
        Given que accedi a la pagina principal
        When  hago el login con "jonathans@avalith.net" y "123jonathans"
        Then estoy autenticado con éxito

    @invalidpass
    Scenario: Contraseña incorrecta
        Given que accedi a la pagina principal
        When  hago el login con "jonathans@avalith.net" y "callefalsa123"
        Then debo ver el siguiente mensaje "El inicio de sesión no es válido."

    @invalidemail
    Scenario: Email inválido
        Given que accedi a la pagina principal
        When  hago el login con "rucurucu@avalith.net" y "123jonathans"
        Then debo ver el siguiente mensaje "El inicio de sesión no es válido."