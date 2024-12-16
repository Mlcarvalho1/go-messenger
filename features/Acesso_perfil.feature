Feature: Acesso do perfil pessoal
As a Usuário
I want to Visualizar as informações do meu perfil 
So that Eu possa revisar e confirmar minhas informações pessoais

Scenario: Visualizar informações do perfil pessoal com sucesso
Given o usuário “Victor Mendonça” está na página inicial
When o usuário aciona o botão “Meu Perfil”
Then o usuário é redirecionado para a página “Meu Perfil”
And visualiza as informações pessoais, incluindo: nome do usuário, foto de perfil, opção de excluir conta
