```markdown
# 🔐 GO_AUTH_JWT – API d’authentification avec JWT en Go

Ce projet est une API RESTful permettant l’**inscription**, la **connexion** et l’**accès à une route protégée** à l’aide de **JSON Web Tokens (JWT)**. Elle est développée avec **Go**, le framework **Gin**, et utilise **GORM** pour l’accès à la base de données.

---

## 🎯 Fonctionnalités principales

- ✅ Inscription d’un utilisateur (`/Signup`)
- ✅ Connexion et génération de JWT (`/Login`)
- ✅ Stockage du token dans un cookie HTTP-only
- ✅ Middleware `RequireAuth` pour sécuriser les routes (`/Validate`)
- ✅ Hachage sécurisé des mots de passe (`bcrypt`)
- ✅ Validation de token et de son expiration

---

## 📁 Structure du projet

```

GO\_AUTH\_JWT/
├── controllers/       # Logique des endpoints : Signup, Login, Validate
├── middleware/        # Middleware pour protéger les routes avec JWT
├── models/            # Modèle User
├── initializers/      # Connexion DB & chargement .env
├── main.go            # Point d’entrée principal
├── .env               # Variables d’environnement (non versionné)
├── go.mod             # Dépendances Go

````

---

## ⚙️ Installation et exécution

### 1. Cloner le projet

```bash
git clone https://github.com/TON-UTILISATEUR/GO_AUTH_JWT.git
cd GO_AUTH_JWT
````

### 2. Configurer les variables d’environnement

Crée un fichier `.env` :

```env
DB_URL=postgres://user:password@localhost:5432/dbname
SECRET=ma_clé_secrète_jwt
```

### 3. Installer les dépendances

```bash
go mod tidy
```

### 4. Lancer le serveur

```bash
go run main.go
```

L'API sera disponible sur :
📍 `http://localhost:8080`

---

## 🧪 Endpoints disponibles

| Méthode | URL         | Description                         |
| ------- | ----------- | ----------------------------------- |
| POST    | `/Signup`   | Créer un nouvel utilisateur         |
| POST    | `/Login`    | Connexion et retour d’un JWT cookie |
| GET     | `/Validate` | Route protégée par JWT (auth)       |

---

## 🔐 Authentification

* Le token est généré après `/Login`
* Il est stocké dans un **cookie sécurisé** (`Authorization`)
* Le middleware `RequireAuth` vérifie le token JWT et l’utilisateur lié

---

## 📦 Technologies utilisées

* [Go](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/)
* [JWT](https://github.com/golang-jwt/jwt)
* [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
* [gotenv](https://github.com/subosito/gotenv)

---

## 📄 Exemple de réponse à la connexion

```http
POST /Login
{
  "Email": "test@example.com",
  "Password": "123456"
}
```

**Réponse :**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "message": "Connexion réussie"
}
```

> 🚀 Ce projet peut servir de base à tout système d’authentification avec JWT dans vos applications Go !
