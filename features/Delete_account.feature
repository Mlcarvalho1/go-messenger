Feature: Exclusão de conta
As a Usuário
I want to Excluir minha conta
So that Eu não tenha minha conta exibida para outros usuários

Scenario: Abrir modal de confirmação para exclusão de conta
Given o usuário “Victor Mendonça” está na página “Meu Perfil”
When o usuário “Victor Mendonça” aciona o botão “Excluir conta”
Then o modal de confirmação é exibido
And o modal contém a mensagem “Você está certo disso?”
And o modal contém os botões “Sim” e “Não”

Scenario: Confirmar exclusão de conta
Given o modal de confirmação está aberto
When o usuário aciona o botão “Sim” 
Then a conta é excluída
And o usuário recebe a mensagem “Conta excluída com sucesso”
