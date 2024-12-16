Feature: Recuperação de senha            
  As a: Usuário
  I want to: Recuperar minha senha
  So that: Eu possa logar na minha conta com sucesso

Scenario: Acionamento de “Resetar senha”
  Given o usuário “João” está na tela “Esqueci minha senha”
  When o usuário “João” preencher o e-mail
  And acionar o “botão” “Resetar senha”
  Then o usuário “João” recebe uma mensagem indicando que um link de redefinição de senha foi enviado para o seu e-mail
  And o link é enviado caso o e-mail esteja cadastrado na aplicação

Scenario: Acionamento de “Resetar senha” sem preencher o e-mail
  Given o usuário “João” está na tela “Esqueci minha senha”
  When o usuário “João” deixa o campo “e-mail” vazio
  And aciona o “botão” “Resetar Senha”
  Then o usuário “João” recebe uma mensagem de erro “Todos os campos devem ser preenchidos”

Scenario: Acionamento de “Login”
  Given o usuário “João” está na tela “Esqueci minha senha”
  When o usuário “João” acionar o “botão” “Login”
  Then o usuário “João” é redirecionado para a tela “Login”
