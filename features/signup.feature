    Feature : Cadastro de usuário
    As a usuário não cadastrado
    I want to me cadastrar no sistema
    So that eu possa enviar e receber mensagem


    Scenario: Cadastro de usuário com sucesso
        Given o usuário está na página de “cadastro de usuário”
        When o usuário preenche os campos "nome" com "Rafael Alves", "email" com "example@email.com", "senha" com "senhaSegura123." , "confirmar senha" com "senhaSegura123." e envia uma foto de perfil
        And aciona o botão “criar conta”
        Then o usuário está na página de “Confirme o seu email”
        When o usuário confirma o email
        Then o usuário está na página de “login”

    Scenario: Cadastro de usuário mal sucedido ( email inválido )
        Given o usuário está na página de “cadastro de usuário”
        When o usuário preenche os campos “nome” com “Rafael Alves”, “email” com “exampleemail” e “senha” com  “senhaSegura123.” , "confirmar senha" com "senhaSegura123." e envia uma foto de perfil
        And clica no botão “criar conta”
        Then o sistema exibe uma mensagem de erro "Email inválido. Por favor, insira um email válido."