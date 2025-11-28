-- ==================== SCRIPTS SQL - Base de Datos go_fullstack ====================
-- Este archivo contiene scripts útiles para la gestión de la base de datos
-- Autor: Rami
-- Fecha: Noviembre 2025

-- ==================== CREAR BASE DE DATOS ====================
-- Ejecutar si la base de datos no existe

CREATE DATABASE IF NOT EXISTS go_fullstack
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE go_fullstack;

-- ==================== INSERTAR DATOS INICIALES ====================

-- Estados (Activo/Inactivo) - IMPORTANTE: Ejecutar después de la primera migración
INSERT INTO estados (nombre)
VALUES ('Activo'), ('Inactivo')
ON DUPLICATE KEY UPDATE nombre = VALUES(nombre);

-- ==================== DATOS DE PRUEBA - CATEGORÍAS ====================

INSERT INTO categoria (nombre, slug, created_at, updated_at)
VALUES
    ('Bebidas', 'bebidas', NOW(), NOW()),
    ('Sopas', 'sopas', NOW(), NOW()),
    ('Tragos y cócteles', 'tragos-y-cocteles', NOW(), NOW()),
    ('Postres', 'postres', NOW(), NOW()),
    ('Platos principales', 'platos-principales', NOW(), NOW()),
    ('Entradas', 'entradas', NOW(), NOW())
ON DUPLICATE KEY UPDATE nombre = VALUES(nombre);

-- ==================== DATOS DE PRUEBA - RECETAS ====================
-- NOTA: Ajusta los IDs de categoria_id y usuario_id según tus datos

INSERT INTO receta (categoria_id, usuario_id, nombre, slug, tiempo, foto, descripcion, fecha, created_at, updated_at)
VALUES
    (1, 1, 'Limonada fresca', 'limonada-fresca', '10 min', 'limonada.jpg', 'Bebida refrescante de limón con hielo y menta.', NOW(), NOW(), NOW()),
    (2, 1, 'Sopa de tomate', 'sopa-de-tomate', '30 min', 'sopa_tomate.jpg', 'Sopa cremosa de tomate con albahaca.', NOW(), NOW(), NOW()),
    (3, 1, 'Cóctel Margarita', 'coctel-margarita', '7 min', 'margarita.jpg', 'Trago clásico con tequila, triple sec y limón.', NOW(), NOW(), NOW()),
    (4, 1, 'Brownie de chocolate', 'brownie-de-chocolate', '45 min', 'brownie.jpg', 'Postre de chocolate intenso con nueces.', NOW(), NOW(), NOW()),
    (5, 1, 'Pasta a la carbonara', 'pasta-a-la-carbonara', '25 min', 'carbonara.jpg', 'Pasta italiana con salsa de huevo, queso y panceta.', NOW(), NOW(), NOW())
ON DUPLICATE KEY UPDATE nombre = VALUES(nombre);

-- ==================== CONSULTAS ÚTILES ====================

-- Ver todas las categorías (incluyendo las eliminadas)
SELECT * FROM categoria;

-- Ver solo categorías activas (no eliminadas)
SELECT * FROM categoria WHERE deleted_at IS NULL;

-- Ver recetas con su categoría y usuario
SELECT
    r.id,
    r.nombre AS receta,
    c.nombre AS categoria,
    u.nombre AS usuario,
    r.tiempo,
    r.fecha
FROM receta r
INNER JOIN categoria c ON r.categoria_id = c.id
INNER JOIN usuario u ON r.usuario_id = u.id
WHERE r.deleted_at IS NULL
ORDER BY r.fecha DESC;

-- Contar recetas por categoría
SELECT
    c.nombre AS categoria,
    COUNT(r.id) AS total_recetas
FROM categoria c
LEFT JOIN receta r ON c.id = r.categoria_id AND r.deleted_at IS NULL
WHERE c.deleted_at IS NULL
GROUP BY c.id, c.nombre
ORDER BY total_recetas DESC;

-- Ver usuarios registrados y su estado
SELECT
    u.id,
    u.nombre,
    u.correo,
    e.nombre AS estado,
    u.fecha
FROM usuario u
INNER JOIN estado e ON u.estado_id = e.id
ORDER BY u.fecha DESC;

-- Ver mensajes de contacto recientes
SELECT
    nombre,
    correo,
    telefono,
    LEFT(mensaje, 50) AS mensaje_corto,
    fecha
FROM contacto
ORDER BY fecha DESC
LIMIT 10;

-- ==================== MANTENIMIENTO ====================

-- Limpiar registros marcados como eliminados (soft delete)
-- ⚠️ CUIDADO: Esto elimina permanentemente los registros

-- Eliminar categorías eliminadas hace más de 30 días
DELETE FROM categoria
WHERE deleted_at IS NOT NULL
AND deleted_at < DATE_SUB(NOW(), INTERVAL 30 DAY);

-- Eliminar recetas eliminadas hace más de 30 días
DELETE FROM receta
WHERE deleted_at IS NOT NULL
AND deleted_at < DATE_SUB(NOW(), INTERVAL 30 DAY);

-- ==================== RESTAURAR SOFT DELETES ====================

-- Restaurar una categoría específica
UPDATE categoria
SET deleted_at = NULL
WHERE id = 1 AND deleted_at IS NOT NULL;

-- Restaurar una receta específica
UPDATE receta
SET deleted_at = NULL
WHERE id = 1 AND deleted_at IS NOT NULL;

-- ==================== ESTADÍSTICAS ====================

-- Total de registros por tabla
SELECT
    'Categorías' AS tabla,
    COUNT(*) AS total,
    SUM(CASE WHEN deleted_at IS NULL THEN 1 ELSE 0 END) AS activos,
    SUM(CASE WHEN deleted_at IS NOT NULL THEN 1 ELSE 0 END) AS eliminados
FROM categoria
UNION ALL
SELECT
    'Recetas',
    COUNT(*),
    SUM(CASE WHEN deleted_at IS NULL THEN 1 ELSE 0 END),
    SUM(CASE WHEN deleted_at IS NOT NULL THEN 1 ELSE 0 END)
FROM receta
UNION ALL
SELECT
    'Usuarios',
    COUNT(*),
    SUM(CASE WHEN estado_id = 1 THEN 1 ELSE 0 END),
    SUM(CASE WHEN estado_id = 2 THEN 1 ELSE 0 END)
FROM usuario
UNION ALL
SELECT
    'Contactos',
    COUNT(*),
    NULL,
    NULL
FROM contacto;

-- ==================== RESPALDO Y RESTAURACIÓN ====================

-- Exportar base de datos (ejecutar en terminal, no en MySQL)
-- mysqldump -u root -p go_fullstack > backup_go_fullstack.sql

-- Importar base de datos (ejecutar en terminal, no en MySQL)
-- mysql -u root -p go_fullstack < backup_go_fullstack.sql

-- ==================== ÍNDICES ADICIONALES (OPCIONAL) ====================
-- Estos índices pueden mejorar el rendimiento en consultas frecuentes

-- Índice en nombre de recetas (para búsquedas)
CREATE INDEX idx_receta_nombre ON receta(nombre);

-- Índice en slug de recetas (para búsquedas por URL)
CREATE INDEX idx_receta_slug ON receta(slug);

-- Índice en correo de usuarios (para login)
CREATE INDEX idx_usuario_correo ON usuario(correo);

-- Índice en categoria_id de recetas (para filtros por categoría)
CREATE INDEX idx_receta_categoria ON receta(categoria_id);

-- Índice en usuario_id de recetas (para filtros por usuario)
CREATE INDEX idx_receta_usuario ON receta(usuario_id);

-- ==================== LIMPIAR ÍNDICES (SI ES NECESARIO) ====================

-- DROP INDEX idx_receta_nombre ON receta;
-- DROP INDEX idx_receta_slug ON receta;
-- DROP INDEX idx_usuario_correo ON usuario;
-- DROP INDEX idx_receta_categoria ON receta;
-- DROP INDEX idx_receta_usuario ON receta;

-- ==================== VERIFICAR ESTRUCTURA ====================

-- Ver estructura de una tabla
DESCRIBE categoria;
DESCRIBE receta;
DESCRIBE usuario;
DESCRIBE contacto;
DESCRIBE estado;

-- Ver todos los índices de una tabla
SHOW INDEX FROM receta;

-- Ver el tamaño de las tablas
SELECT
    table_name AS tabla,
    ROUND(((data_length + index_length) / 1024 / 1024), 2) AS tamaño_mb
FROM information_schema.TABLES
WHERE table_schema = 'go_fullstack'
ORDER BY (data_length + index_length) DESC;

-- ==================== FIN ====================

