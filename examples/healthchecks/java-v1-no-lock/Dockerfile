FROM openjdk:latest

RUN mkdir -p /src /app
WORKDIR /app
COPY Main.java /src/Main.java

RUN javac /src/Main.java -d /src/
ENTRYPOINT ["java", "-classpath", "/src/", "Main"]
