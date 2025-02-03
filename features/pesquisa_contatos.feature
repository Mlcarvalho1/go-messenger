Essa feacture tem o objetivo de pesquisar na lista de contatos por um contato especifico. 

Cenário 1: Pesquisa de usuário na lista de contatos
Scenario: Pesquisa por nome na lista de contatos
Given: O usuário “Luan Romancini” está na página “lista de contatos”
And: ve a aba de pesquisa “pesquisar por contatos”
When: pesquisa pelo nome "João" na barra de pesquisa
Then: vê todos os usuários cujo nome contém a string "João"
Then: novamente incluindo Then apenas para fazer o exercício. 

Cenário 2: Abrir conversa com usuário da lista de contatos que nunca troquei mensagem
Scenario: Abrir conversa com usuário da lista de contatos
Given: O usuário “Luan Romancini” esta na pagina “lista de contatos”
And: ve o perfil do "João" aberto
And: ve o botão "abrir conversa" ao lado do perfil
When: clica no botão "abrir conversa"
Then: Uma aba de conversa é aberta com “João”
And: ve uma conversa vazia com “João”

