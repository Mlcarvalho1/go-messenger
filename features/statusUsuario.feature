Feature: Status do usuário         
  As a: Usuário
  I want to: Saber se outro usuário está online ou a última vez que esteve online
  So that: Eu possa saber a disponibilidade dele para responder 

Scenario: Exibir que um usuário está online
  Given o usuário “João” está na tela “conversas”
  And o usuário “Rafael” está online
  When o usuário “João” abrir a tela “chat com Rafael” 
  Then o usuário “João” vê uma indicação visual de que o usuário “Rafael” está online

Scenario: Exibir a última vez que um usuário esteve online
  Given o usuário “João” está na tela “conversas”
  And o usuário “Rafael” está desconectado
  When o usuário “João” abrir a tela “chat com Rafael” 
  Then o usuário “João” vê uma indicação visual do último momento em que o usuário “Rafael” esteve online
