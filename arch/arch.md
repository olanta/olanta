
Agora vamos nos ater ao core que segue esta estrutura aqui:

core/
    cmd/
        olanta/
            main.go
    config/
        dist.go
        config.go
        const.go
    internal/
        analyzer/
            analyzer.go
        helpers/
            logger.go
            messages.go
        indexer/
            indexer.go
            config.go
        models/
            models.go
        server/
            server.go
        utils/
            file.go
            yaml_loader.go
    build/
        buildfile.go
        main.go




scanner/
    cmd/
        scanner/
            main.go
    internal/
        rules/
            java_rules.yaml
            python_rules.yaml
        scanner/
            scanner.go
            java_scanner.go
            python_scanner.go
        utils/
            yaml_loader.go
    build/
        buildfile.go
        main.go
.gitignore
go.mod
README.md


