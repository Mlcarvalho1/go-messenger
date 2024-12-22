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

    Scenario: Cadastro com email já cadastrado
        Given o usuário está na página de “cadastro de usuário”
        And já existe um usuário cadastrado com o email “example@email.com” no sistema
        When o usuário preenche os campos “nome” com “Rafael Alves”, “email” com “example@email.com” e “senha” com  “senhaSegura123.”, "confirmar senha" com "senhaSegura123." e envia uma foto de perfil
        And clica no botão “criar conta”
        Then o sistema exibe uma mensagem de erro "O email informado já está em uso"

    Scenario: Cadastro com senha Fraca
        Given o usuário está na página de “cadastro de usuário”
        When o usuário preenche os campos “nome” com “Rafael Alves”, “email” com “example@email.com” e “senha” com  “1234” , "confirmar senha" com "1234." e envia uma foto de perfil
        And clica no botão “criar conta”
        Then o sistema exibe uma mensagem de erro "Senha não cumpre os requisitos de segurança."

    Scenario: Cadastro com senha e senha de confirmação diferentes
        Given o usuário está na página de “cadastro de usuário”
        When o usuário preenche os campos “nome” com “Rafael Alves”, “email” com “example@email.com” e “senha” com  “senhaSegura12.” , "confirmar senha" com "snhaSegura1." e envia uma foto de perfil
        And clica no botão “criar conta”
        Then o sistema exibe uma mensagem de erro "Senha não cumpre os requisitos de segurança."

    Scenario: Cadastro sem preencher todos os campos
        Given o usuário está na página de “cadastro de usuário”
        When o usuário preenche os campos “email” com “example@email.com” e “senha” com  “senhaSegura12.” , "confirmar senha" com "senhaSegura12.." e envia uma foto de perfil
        And clica no botão “criar conta”
        Then o sistema exibe uma mensagem de erro "Todos os campos devem ser preenchidos."