CREATE TABLE IF NOT EXISTS preguntas (
    id SERIAL PRIMARY KEY,
    pregunta TEXT NOT NULL,
    respuesta TEXT NOT NULL
);