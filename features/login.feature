Feature: Login            
  As a: Usuário
  I want to: Logar na minha conta
  So that: Eu possa utilizar as funcionalidades do site

Scenario: Login feito com sucesso
  Given o usuário “João” está na página “Login”
  And já possui uma conta no sistema
  When o usuário “João” preenche corretamente seu “Login” e “Senha”
  Then o usuário “João” acessa a sua conta com sucesso
  And vai para a página inicial

Scenario: Login feito incorretamente
  Given o usuário “João” está na página “Login”
  And já possui uma conta no sistema
  When o usuário “João” preenche incorretamente a sua “senha”
  Then o usuário “João” recebe uma mensagem de e-mail/senha inválido

Scenario: Login feito incorretamente
  Given o usuário “João” está na página “Login”
  And não já possui uma conta no sistema 
  When o usuário “João” preenche o “login” e “senha”
  Then o usuário “João” recebe uma mensagem de e-mail/senha inválido
