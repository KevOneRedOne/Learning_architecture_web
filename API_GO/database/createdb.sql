-- Database: m1_myapp_archi_log

-- DROP DATABASE IF EXISTS m1_myapp_archi_log;

CREATE ROLE postgres;

CREATE DATABASE m1_myapp_archi_log
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

GRANT ALL ON DATABASE m1_myapp_archi_log TO postgres;

CREATE TABLE IF NOT EXISTS public.articles(
    id serial PRIMARY KEY NOT NULL,
    title text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    date text COLLATE pg_catalog."default"
);

ALTER TABLE IF EXISTS public.articles
    OWNER to postgres;

INSERT INTO articles(title,description,date) VALUES ('Les Avantages du Cloud Computing pour les Entreprises','Cet article explore les avantages du cloud computing, y compris la flexibilité, l''évolutivité et la réduction des coûts. Il examine également les meilleures pratiques pour migrer vers le cloud.','2023-05-12');
INSERT INTO articles(title,description,date) VALUES ('Sécurité dans le Cloud : Mythes et Réalités','Cet article dissipe les mythes courants entourant la sécurité du cloud et propose des conseils pratiques pour assurer une protection adéquate des données dans le cloud.','2023-06-03');
INSERT INTO articles(title,description,date) VALUES ('Les Différents Modèles de Service Cloud Expliqués','Découvrez les principaux modèles de service cloud : IaaS, PaaS et SaaS. Apprenez comment chacun fonctionne et quelles sont les applications les plus adaptées à chaque modèle.','2023-07-18');
INSERT INTO articles(title,description,date) VALUES ('Le Cloud Hybride : Un Pont entre Cloud Public et Privé','Cet article explore les avantages du cloud hybride, qui combine les infrastructures cloud publiques et privées. Il examine également les considérations importantes lors de la mise en place d''une architecture hybride.','2023-08-05');
INSERT INTO articles(title,description,date) VALUES ('Comment Optimiser les Coûts dans le Cloud','Découvrez des stratégies et des outils pour optimiser les coûts associés à l''utilisation du cloud. Cela inclut la gestion des ressources, le suivi des dépenses et l''utilisation de modèles d''achat flexibles.','2023-09-22');
INSERT INTO articles(title,description,date) VALUES ('L''Impact du Cloud Computing sur l''Innovation','Cet article examine comment le cloud computing a révolutionné l''innovation dans les entreprises. Il présente des études de cas et des exemples concrets de succès grâce à l''adoption du cloud.','2023-10-02');
INSERT INTO articles(title,description,date) VALUES ('Migrer vers le Cloud : Étapes Cruciales à Suivre','Découvrez les étapes essentielles à suivre lors de la migration vers le cloud. Cela inclut l''évaluation des besoins, la planification de la migration et la gestion des risques.','2023-10-11');
INSERT INTO articles(title,description,date) VALUES ('Le Rôle Clé de DevOps dans le Cloud Computing','Cet article explore l''importance de DevOps dans le contexte du cloud computing. Il met en lumière les meilleures pratiques pour intégrer DevOps dans les opérations cloud.','2023-10-15');
INSERT INTO articles(title,description,date) VALUES ('L''Évolution des Technologies de Stockage dans le Cloud','Découvrez comment les technologies de stockage évoluent dans le cloud pour répondre aux besoins croissants en matière de performance et de capacité. Explorez les tendances actuelles en matière de stockage cloud.','2023-10-19');
INSERT INTO articles(title,description,date) VALUES ('L''Intelligence Artificielle dans le Cloud : Opportunités et Défis','Cet article examine comment l''intelligence artificielle s''intègre dans les services cloud et présente les avantages et les défis de cette convergence.','2023-10-25');







