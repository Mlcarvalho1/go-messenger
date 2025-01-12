Feature: pesquisa de Chats
As a um Usuário
I want to pesquisar chats por nome na lista de chats
So that eu possa rapidamente encontrar o chat que estou procurando

Scenario: Pesquisa por nome na lista de chats com correspondência
Given: o usuário "Ian Melo" está na página "lista de chats"
And: existem chats que contenham a string "João"
And: vê a aba de pesquisa "pesquisar chats"
When: pesquisa pelo nome "João" na barra de pesquisa
Then: vê todos os chats cujo nome contém a string "João" em ordem de mensagem mais recente

Scenario: pesquisa por nome na lista de chats sem correspondência
Given: o usuário "Ian Melo" está na página "lista de chats"
And: não existem chats que contenham a string "Henricão"
When: pesquisa pelo nome "Henricão" na barra de pesquisa
Then: vê a mensagem "não foram encontrados chats"
