# Contributing — OpenDev Hub

Bienvenue sur OpenDev Hub, la plateforme centrale d'OpenDev Mada.

Ce fichier explique comment contribuer, ce qui reste à construire, et comment soumettre ton travail proprement. Lis-le en entier avant d'ouvrir une PR.

---

## C'est quoi OpenDev Hub ?

OpenDev Hub est une plateforme web pour centraliser toute l'activité de la communauté OpenDev Mada : profils membres, défis hebdomadaires, projets open source, système de points et classement.

**Stack :**
- Backend : Go 1.25, Fiber v2, PostgreSQL, sqlx
- Frontend : React
- Auth : JWT, bcrypt

---

## Structure du projet

```
OpenDev/
├── cmd/
│   └── api/
│       └── main.go           # Point d'entrée — lance le serveur
├── internal/
│   ├── server/
│   │   ├── server.go         # Configuration Fiber (middlewares, error handler)
│   │   └── routes.go         # Toutes les routes HTTP
│   ├── database/
│   │   └── database.go       # Connexion PostgreSQL
│   ├── users/
│   │   ├── user.go           # Modèle User
│   │   └── repository.go     # Requêtes SQL Users
│   └── auth/                 # À construire — logique JWT
├── migrations/
│   ├── 001_create_user.sql   # Table users — déjà définie
│   ├── 002_create_project.sql # Table projects — à compléter
│   └── 003_create_challenge.sql # Table challenges — à compléter
├── Maquette/
│   └── opendevmada.html      # Maquette frontend de référence
├── .env.example              # Variables d'environnement à copier
└── go.mod
```

---

## Lancer le projet en local

**Prérequis :** Go 1.21+, PostgreSQL, Git

```bash
# 1. Cloner le repo
git clone https://github.com/Safidy-Mahefa/OpenDev-hub.git
cd OpenDev-hub

# 2. Copier le fichier env
cp .env.example .env
# Remplir DATABASE_URL, PORT, JWT_SECRET dans .env

# 3. Créer la base de données PostgreSQL
createdb opendev_hub

# 4. Exécuter les migrations
psql -d opendev_hub -f migrations/001_create_user.sql
psql -d opendev_hub -f migrations/002_create_project.sql #encore vide
psql -d opendev_hub -f migrations/003_create_challenge.sql #encore vide

# 5. Installer les dépendances
go mod tidy

# 6. Lancer le serveur
go run cmd/api/main.go

# Le serveur tourne sur http://localhost:3000
# Tester : GET http://localhost:3000/health
```

---

## Ce qui reste à construire

Chaque tâche ci-dessous est indépendante. Tu prends une tâche, tu crées une branche, tu fais une PR. Pas besoin de tout comprendre pour commencer.

---

### BACKEND — Go / Fiber / PostgreSQL

#### Étape 1 — Authentification
> Dossier : `internal/auth/`

- [ ] `POST /auth/register` — inscription (username, email, password) — hash avec bcrypt
- [ ] `POST /auth/login` — vérifier les credentials, retourner un JWT
- [ ] Middleware JWT — protéger les routes privées
- [ ] `POST /auth/logout` — invalider le token côté client

Dépendances Go à installer :
```bash
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v5
```

---

#### Étape 2 — Users CRUD
> Dossier : `internal/users/`

Routes à ajouter dans `routes.go` :

- [ ] `GET /users` — liste tous les membres (public)
- [ ] `GET /users/:id` — profil d'un membre (public)
- [ ] `PUT /users/:id` — modifier son profil — protégé JWT (bio, ville, githubUsername, linkedinUrl, portfolioUrl, avatarUrl)
- [ ] `DELETE /users/:id` — supprimer son compte — protégé JWT

Le modèle User est déjà défini dans `internal/users/user.go`. Les fonctions `GetAll` et `Create` existent dans `repository.go` — s'en inspirer pour les nouvelles fonctions.

---

#### Étape 3 — Projets
> Dossier : `internal/projects/` — à créer

**Migration à compléter** (`migrations/002_create_project.sql`) :
```sql
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    stack VARCHAR(255),
    status VARCHAR(50) DEFAULT 'open',  -- open / active / paused / delivered
    author_id VARCHAR(10) REFERENCES users(id),
    github_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Routes à implémenter :
- [ ] `POST /projects` — soumettre un projet — protégé JWT
- [ ] `GET /projects` — liste tous les projets (filtrable par status)
- [ ] `GET /projects/:id` — détail d'un projet
- [ ] `PUT /projects/:id` — modifier un projet — protégé JWT (auteur uniquement)
- [ ] `POST /projects/:id/join` — rejoindre un projet — protégé JWT
- [ ] `GET /projects/:id/contributors` — liste des contributeurs

---

#### Étape 4 — Défis
> Dossier : `internal/challenges/` — à créer

**Migration à compléter** (`migrations/003_create_challenge.sql`) :
```sql
CREATE TABLE challenges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    stack VARCHAR(255),
    points INT DEFAULT 20,
    deadline TIMESTAMP,
    status VARCHAR(50) DEFAULT 'open',  -- open / closed
    created_by VARCHAR(10) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    challenge_id UUID REFERENCES challenges(id),
    user_id VARCHAR(10) REFERENCES users(id),
    github_url VARCHAR(255) NOT NULL,
    note TEXT,
    score INT,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Routes à implémenter :
- [ ] `POST /challenges` — créer un défi — admin seulement
- [ ] `GET /challenges` — liste tous les défis
- [ ] `GET /challenges/:id` — détail d'un défi
- [ ] `POST /challenges/:id/submit` — soumettre une solution — protégé JWT
- [ ] `GET /challenges/:id/submissions` — voir toutes les soumissions

---

#### Étape 5 — Points & Classement
> Modifier `internal/users/repository.go`

- [ ] Fonction `AddPoints(db, userID, points)` — ajoute des points à un user
- [ ] Appelée automatiquement quand une soumission est notée
- [ ] `GET /leaderboard` — classement par `totalPoints` DESC
- [ ] `GET /leaderboard/season` — classement par `seasonPoints` DESC

---

### FRONTEND — React

> La maquette HTML de référence est dans la branche `maquette-frontend`. C'est la base visuelle à suivre.

#### Étape 1 — Setup
- [ ] Initialiser le projet React dans un dossier `/frontend`
- [ ] Configurer Axios pour les appels API (`baseURL: http://localhost:3000`)
- [ ] Mettre en place React Router v6
- [ ] Créer les composants de base : `Navbar`, `Footer`, `PrivateRoute`
- [ ] Gérer le token JWT dans le contexte React (AuthContext)

---

#### Étape 2 — Auth
- [ ] Page `/register` — formulaire inscription
- [ ] Page `/login` — formulaire connexion
- [ ] Stockage du JWT dans localStorage
- [ ] Redirection automatique si connecté / déconnecté
- [ ] Bouton déconnexion dans la Navbar

---

#### Étape 3 — Pages principales
- [ ] Page `/` — accueil — présentation de la communauté + stats (membres, projets, défis)
- [ ] Page `/members` — liste des membres avec filtre par stack ou ville
- [ ] Page `/members/:id` — profil public d'un membre (bio, stack, points, GitHub, projets)
- [ ] Page `/challenges` — liste des défis actifs et passés
- [ ] Page `/challenges/:id` — détail d'un défi + formulaire de soumission
- [ ] Page `/projects` — liste des projets ouverts aux contributions
- [ ] Page `/projects/:id` — détail d'un projet + liste des contributeurs
- [ ] Page `/leaderboard` — classement global et saisonnier

---

#### Étape 4 — Dashboard membre
> Route : `/dashboard` — protégée, connecté uniquement

- [ ] Modifier son profil (bio, ville, GitHub, LinkedIn, avatar)
- [ ] Voir ses soumissions de défis et leurs scores
- [ ] Voir les projets auxquels on contribue
- [ ] Voir ses points totaux et saisonniers

---

## Comment soumettre une contribution

### 1. Fork ou clone le repo
```bash
git clone https://github.com/Safidy-Mahefa/OpenDev-hub.git
cd OpenDev-hub
```

### 2. Créer une branche avec un nom explicite
```bash
# Format : type/nom-de-la-tâche
git checkout -b feature/auth-register
git checkout -b feature/users-crud
git checkout -b fix/migration-projects
git checkout -b frontend/page-login
```

### 3. Coder, tester, commiter
```bash
git add .
git commit -m "feat: add POST /auth/register with bcrypt"
```

Convention des commits :
- `feat:` — nouvelle fonctionnalité
- `fix:` — correction de bug
- `refactor:` — refactoring sans nouvelle feature
- `docs:` — documentation
- `migration:` — modification de migration SQL

### 4. Ouvrir une Pull Request
- Titre clair : `feat: POST /auth/register`
- Description : ce que tu as fait, comment tester, screenshots si frontend
- Une tâche par PR — pas de PR avec 5 fonctionnalités mélangées

---

## Règles

- Ne jamais commiter le fichier `.env` — utiliser `.env.example`
- Tester ta route ou ta page avant d'ouvrir une PR
- Une PR = une tâche = une branche
- Si tu bloques, poste ta question dans `#entraide` sur Discord

---

## Questions

Discord : serveur OpenDev Mada — canal `#contributions`
Repo : https://github.com/Safidy-Mahefa/OpenDev-hub.git
:)