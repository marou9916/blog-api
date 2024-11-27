# Blog API

Une API simple pour gérer un blog avec des fonctionnalités CRUD (Créer, Lire, Mettre à jour, Supprimer) pour des articles. Ce projet est conçu pour renforcer les compétences en création d'APIs robustes en utilisant **Golang**, avec un accent sur la validation des données, la documentation des endpoints, et les bonnes pratiques de développement.

---

## Table des matières

1. [Introduction](#introduction)
2. [Objectifs du projet](#objectifs-du-projet)
3. [Fonctionnalités](#fonctionnalités)
4. [Technologies utilisées](#technologies-utilisées)
5. [Installation et exécution](#installation-et-exécution)
6. [Structure du projet](#structure-du-projet)
7. [Documentation API](#documentation-api)
8. [Compétences renforcées](#compétences-renforcées)
9. [Améliorations futures](#améliorations-futures)
10. [Contributions](#contributions)
11. [Licence](#licence)

---

## Introduction

Ce projet est une API RESTful développée en **Golang**, conçue pour gérer un blog avec des articles. L'objectif principal est d'offrir une API fiable, bien structurée, et maintenable, tout en explorant des outils modernes pour le développement d'API.

---

## Objectifs du projet

Les objectifs principaux de ce projet sont :

1. **Renforcement des compétences en développement d'API robuste**  
   Appliquer des concepts essentiels tels que la gestion des erreurs, l'organisation du code, et les tests.
   
2. **Validation des données entrantes**  
   Implémenter des mécanismes pour s'assurer que seules des données valides sont traitées par l'API, avec des retours d'erreur clairs en cas de problème.

3. **Documentation d'API**  
   Utilisation d'outils tels que **Swagger** et/ou **Postman** pour documenter et tester les endpoints, facilitant ainsi leur utilisation et leur maintenance.

4. **Mise en œuvre des bonnes pratiques**  
   Respecter des standards tels que REST et des principes de conception comme l'organisation claire du code.

5. **Apprentissage et intégration de bibliothèques clés**  
   Découverte et intégration de bibliothèques comme **Gorm** (ORM pour Golang), **Gin** (framework HTTP), et d'autres outils pertinents.

---

## Fonctionnalités

- Création d'articles
- Consultation d'un ou de plusieurs articles
- Mise à jour des articles
- Suppression d'articles
- Validation des données avant traitement
- Documentation interactive avec Swagger

---

## Technologies utilisées

- **Langage** : [Golang](https://golang.org/)
- **Framework HTTP** : [Gin](https://github.com/gin-gonic/gin)
- **ORM** : [Gorm](https://gorm.io/)
- **Base de données** : SQLite
- **Tests des requêtes** : Postman
- **Documentation API** : Swagger

---

## Installation et exécution

### Prérequis

- Golang (version 1.20 ou supérieure recommandée)
- SQLite (inclus dans ce projet)
- Outil de gestion des dépendances (modèle `go.mod` intégré)

### Étapes

1. **Cloner le projet :**

   ```bash
   git clone https://github.com/marou9916/blog-api.git
   cd blog-api

2. **Installer les dépendances :**

   ```bash
   go mod tidy

3. **Exécuter le projet :**

   ```bash
   go run .

4. **Accéder à l'API :**

   L'API sera disponible à l'adresse http://localhost:8080.

5. **Accéder à la documentation Swagger :**

   Ouvrez http://localhost:8080/swagger/index.html dans votre navigateur.


## Structure du projet 
```
.
├── cmd
│   └── main.go            # Point d'entrée de l'application
├── pkg
│   ├── controllers        # Gestionnaires pour les endpoints
│   ├── db                 # Initialisation et gestion de la base de données
│   ├── models             # Définition des structures de données
│   └── routes             # Configuration des routes
├── go.mod                 # Dépendances du projet
├── gorm.db                # Base de données SQLite
├── README.md              # Documentation du projet
```

## Documentation API

La documentation Swagger est générée automatiquement. Voici les étapes pour y accéder :

    Lancer le serveur (voir Installation et exécution).
    Accéder à la page Swagger à l'adresse suivante :
    http://localhost:8080/swagger/index.html.

Swagger permet d'interagir directement avec l'API et d'explorer ses fonctionnalités.

## Compétences renforcées

Ce projet m'a permis de travailler et de renforcer plusieurs compétences clés :

   - **Conception d'API robuste :**
     Apprentissage des bonnes pratiques pour structurer une API RESTful fiable et performante.

    - **Validation des données entrantes :**
     Implémentation de mécanismes pour s'assurer que seules des données valides sont acceptées et traiter les erreurs efficacement.

    - **Documentation des API :**
     Utilisation d'outils comme Swagger pour documenter les endpoints de manière claire et interactive.

    - **Gestion de la base de données :**
     Familiarisation avec Gorm, un ORM puissant pour interagir avec SQLite, et migration automatique des modèles.

    - **Prise en main de Gin :**
     Apprentissage approfondi de Gin, un framework HTTP rapide et léger, pour construire une API RESTful efficace.

    - **Gestion des erreurs :**
     Gestion rigoureuse des erreurs dans les différentes couches de l'application.

## Améliorations futures

    - **Ajouter une gestion des utilisateurs et l'authentification JWT.**

    - **Intégrer une base de données plus robuste comme PostgreSQL.**

    - **Mettre en place des tests unitaires et des tests d'intégration.**

    - **Ajouter un système de pagination pour les listes d'articles.**

    - **Implémenter un déploiement automatisé (CI/CD).**

## Contributions 

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir des issues ou des pull requests pour proposer des améliorations ou signaler des bugs.

## License 

Ce projet est sous licence MIT. Vous êtes libre de l'utiliser, de le modifier et de le distribuer conformément aux termes de cette licence. Voir le fichier LICENSE pour plus d'informations.