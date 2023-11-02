FROM postgres:16.0-alpine3.18

# Copy the database initialization script:
COPY .env /tmp/env.list
COPY ./*.sql /docker-entrypoint-initdb.d/

# Set the environment variables:
ENV POSTGRES_DB=DB_NAME
ENV POSTGRES_USER=DB_USER
ENV POSTGRES_PASSWORD=DB_PASSWORD

RUN chmod a+r /docker-entrypoint-initdb.d/*.sql

# Expose the port:
EXPOSE 5432

# Run the database:
CMD ["docker-entrypoint.sh", "postgres"]




#Docker Build
#docker build -t m1_custom_postgres .

#Docker Run
#docker run -d --name m1_custom_postgres --env-file .env -p 5432:5432 -d m1_custom_postgres