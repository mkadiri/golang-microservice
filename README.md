#### Setup
```    
# Build the application
bash build.sh

# Running the application
docker-compose up -d

# View logging 
docker logs -f {container-name}
```

#### Using the application

to access service to go 

`http://localhost:8181`
    
#### API endpoints


#### DEV mode
Setting `MODE=DEV` will enable dev fixtures to be run (seed data)