const DB = require('../../database')


const CheckUsers=async()=>{
    let username = 'admin'
    let password = '123456'
    await DB.query(
        `
            CREATE TABLE IF NOT EXISTS Users(
                id SERIAL NOT NULL,
                username varchar(30) NOT NULL,
                password varchar(30) NOT NULL,
                PRIMARY KEY (id)
            )
        `
    )
    .then(res=>{
        let sql_stm = `
            SELECT username FROM Users
            WHERE EXISTS
            (SELECT username FROM Users WHERE username='${username}');
        `
        DB.query(sql_stm)
        .then(res2=async=>{
            if(res2.rowCount==0){
                DB.query(
                    `
                        INSERT INTO Users(username, password)
                        VALUES('${username}', '${password}');
                    `
                )
            }
        })
        .catch(err2=>{console.log("Error from insert into Users @ CheckUsers : ",err2)})
    })
    .catch(err=>{
        console.log("Error from CheckUsers : ",err)
    })
}