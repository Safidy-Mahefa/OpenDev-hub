### Structude de dossier du serveur GO:
OpenDev
    cmd/            <-----  Point d'entree du programme / Lancement
        api/
            main.go
    internal/       <----- Logique métier privé (encapsulé par go)
        server/         <----- Point d'entree du serveur HTTP
            server.go
            routes.go
        database/       <----- Connection avec la base de Donnee Postgre SQL
        auth/           <----- Logique Sécurité connexion/inscription...
        users/          <----- Logique métier utilisateur
    mogrations/         <----- Template base de donnees pour les migrations
    .env            <----- Infos confidentiels non commités
    .env.example    <----- Exemple commité