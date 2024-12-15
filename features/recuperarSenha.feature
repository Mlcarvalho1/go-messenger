Feature: Recuperação de senha            
  As a: Usuário
  I want to: Recuperar minha senha
  So that: Eu possa logar na minha conta com sucesso

Scenario: Acionamento de “Esqueci minha senha” com um e-mail cadastrado
  Given o usuário “João” está na tela de login
  And o usuário “João” é cadastrado na aplicação
  And o usuário “João” esqueceu a sua senha
  When o usuário “João” acionar o “botão” “Esqueci minha senha” 
  Then o usuário “João” recebe uma mensagem indicando que um link de confirmação foi enviado para o seu e-mail
  And é enviado um link de confirmação para o seu e-mail

Scenario: Acionamento de “Esqueci minha senha” sem um e-mail cadastrado
  Given o usuário “João” está na tela de login
  And o usuário “João” não é cadastrado na aplicação
  When o usuário “João” acionar o “botão” “Esqueci minha senha” 
  Then o usuário “João” recebe uma mensagem indicando que um link de confirmação foi enviado para o seu e-mail
