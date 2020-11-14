FROM maven:3.5-jdk-8-alpine as builder

# Copy local code to the container image.
WORKDIR /app
COPY pom.xml .
COPY aliyun.xml /tmp/settings.xml
COPY src ./src

# Build a release artifact.
RUN mvn package -s /tmp/settings.xml -DskipTests

# Use AdoptOpenJDK for base image.
# It's important to use OpenJDK 8u191 or above that has container support enabled.
# https://hub.docker.com/r/adoptopenjdk/openjdk8
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM adoptopenjdk/openjdk8:jdk8u202-b08-alpine-slim

# Copy the jar to the production image from the builder stage.
COPY --from=builder /app/target/demo-java-*.jar /app.jar

# Run the web service on container startup.
CMD ["java","-Djava.security.egd=file:/dev/./urandom", "-jar","/app.jar"]