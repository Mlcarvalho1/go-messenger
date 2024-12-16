Feature: Login            
  As a: Usuário
  I want to: Logar na minha conta
  So that: Eu possa utilizar as funcionalidades do site

Scenario: Login feito com sucesso
  Given o usuário “João” está na tela “Login”
  And já possui uma conta no sistema
  When o usuário “João” preenche corretamente seu “Login” e “Senha”
  Then o usuário “João” acessa a sua conta com sucesso
  And vai para a página inicial

Scenario: Login incorreto
  Given o usuário “João” está na tela “Login”
  And já possui uma conta no sistema
  When o usuário “João” preenche incorretamente o seu e-mail ou senha
  Then o usuário “João” recebe uma mensagem de e-mail/senha inválido

Scenario: Tentativa de login sem conta criada
  Given o usuário “João” está na tela “Login”
  And não possui uma conta no sistema 
  When o usuário “João” preenche o “login” e “senha”
  Then o usuário “João” recebe uma mensagem de e-mail/senha inválido


Scenario: Acionamento de “Cadastre-se”
  Given o usuário “João” está na tela “Login”
  When o usuário “João” acionar o “botão” “Cadastre-se”
  Then o usuário “João” é enviado para a tela “cadastro”


Scenario: Acionamento de “Esqueci minha senha”
  Given o usuário “João” está na tela “Login”
  When o usuário “João” acionar o “botão” “Esqueci minha senha”
  Then o usuário “João” é enviado para a tela “Esqueci minha senha”

Scenario: Login com campo não preenchido
  Given o usuário “João” está na tela “Login”
  When o usuário “João” deixa o campo “e-mail” ou “senha” vazio 
  Then o usuário “João” recebe uma mensagem de erro “Todos os campos devem ser preenchidos”

