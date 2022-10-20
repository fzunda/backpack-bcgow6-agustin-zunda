package transactions

/*
Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de
un producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage
para buscar el producto. Para esto:
Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico cuyo nombre
puede ser “Before Update”.
El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a
través de un boolean como se observó en la clase.
Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del
producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización.
También debe validarse que el método Read haya sido ejecutado durante el test.
*/
