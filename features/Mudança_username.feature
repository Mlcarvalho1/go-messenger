Feature: Alteração do nome de usuário
As a Usuário
I want to Editar meu nome de usuário
So that Eu possa ter um nome personalizado e atualizado

Scenario: Edição de nome de usuário
Given o usuário “Victor Mendonça” está na página “Meu Perfil”
When o usuário “Victor Mendonça” aciona o botão “Edit username”
And o usuário “Victor Mendonça” digita “Novo nome”
And o usuário “Victor Mendonça” aciona o botão “salvar”
Then o nome de usuário atualizado é exibido na página “Meu Perfil”
