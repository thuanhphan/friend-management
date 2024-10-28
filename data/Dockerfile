FROM migrate/migrate

# Copy all db files
COPY ./migrations /migrations

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]
<<<<<<< HEAD:db/Dockerfile
<<<<<<< HEAD
<<<<<<< HEAD
CMD ["postgresql://postgres:123456@friend-management-go-postgres.db-1:5432/postgres?sslmode=disable up"]
=======
CMD ["postgresql://postgres:Password123@friend-management-go-postgres.db-1:5432/demo?sslmode=disable up"]
>>>>>>> d88719a (Arrange the layered architecture)
=======
CMD ["postgresql://postgres:123456@friend-management-go-postgres.db-1:5432/postgres?sslmode=disable up"]
>>>>>>> eb24ee2 (FM-3)
=======
CMD ["postgresql://postgres:123456@friend-management-postgres.db-1:5432/postgres?sslmode=disable up"]
>>>>>>> 07f2fdf (apply clean architecture):data/Dockerfile
