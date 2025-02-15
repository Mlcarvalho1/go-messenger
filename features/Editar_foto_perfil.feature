Feature: Edição de foto de perfil
As a Usuário
I want to Adicionar, excluir ou trocar a foto de perfil
So that: Eu possa personalizar minha conta com uma foto que me represente,    
atualizar minha foto atualizada ou removê-la caso não deseje exibi-la 

Scenario: Editar foto do perfil pessoal anexando um arquivo .png ou .jpg
Given o usuário “Victor Mendonça” está na modal “Edição de foto”
When o usuário “Victor Mendonça” aciona o botão “Anexar arquivo”
And o usuário “Victor Mendonça” anexa um arquivo “.jpg”
And o usuário “Victor Mendonça” aciona o botão “salvar”
Then o usuário “Victor Mendonça” visualiza sua nova foto de perfil na página “Meu perfil”

Scenario: Anexando foto .png ou .jpg ao perfil pessoal sem salvar
Given o usuário “Victor Mendonça” está na modal “Edição de foto”
When o usuário “Victor Mendonça” aciona o botão “Anexar arquivo”
And o usuário “Victor Mendonça” anexa um arquivo “.jpg”
And o usuário “Victor Mendonça” aciona o botão “cancelar”
Then o usuário “Victor Mendonça” visualiza sua foto de perfil inalterada na página “Meu perfil”

Scenario: Editar foto do perfil pessoal anexando um arquivo que não seja .png ou .jpg
Given o usuário “Victor Mendonça” está no modal “Edição de foto”
When o usuário “Victor Mendonça” aciona o botão “Anexar arquivo”
And o usuário “Victor Mendonça” anexa um arquivo “.pdf”
Then o usuário “Victor Mendonça” recebe uma mensagem “Arquivo inválido”
And o usuário permanece no modal “Edição de foto”

Scenario: Abrir modal de confirmação para exclusão de foto 
Given o usuário “Victor Mendonça” está no modal “Edição de foto” e tem uma foto de perfil
When o usuário “Victor Mendonça” aciona o botão “Excluir foto”
Then o modal de confirmação é exibido 
And o modal contém a mensagem “Você está certo disso?” 
And o modal contém os botões “Sim” e “Não”

Scenario: Confirmar exclusão de foto de perfil 
Given o modal de confirmação está aberto
When o usuário “Victor Mendonça” aciona o botão “Sim” 
Then a foto de perfil é excluída 
And a foto padrão é exibida na página “Meu Perfil”
