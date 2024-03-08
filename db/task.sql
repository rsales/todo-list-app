-- Define a estrutura do banco de dados
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE
);

-- Adiciona algumas tarefas iniciais
INSERT INTO tasks (title, description, completed) VALUES ('Tarefa 1', 'Descrição da Tarefa 1', FALSE);
INSERT INTO tasks (title, description, completed) VALUES ('Tarefa 2', 'Descrição da Tarefa 2', FALSE);
