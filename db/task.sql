-- Define a estrutura do banco de dados
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT
);

-- Adiciona algumas tarefas iniciais
INSERT INTO tasks (title, description) VALUES ('Tarefa 1', 'Descrição da Tarefa 1');
INSERT INTO tasks (title, description) VALUES ('Tarefa 2', 'Descrição da Tarefa 2');