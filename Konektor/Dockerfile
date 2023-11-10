#See https://aka.ms/customizecontainer to learn how to customize your debug container and how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -y libpng-dev libjpeg-dev curl libxi6 build-essential libgl1-mesa-glx
RUN curl -sL https://deb.nodesource.com/setup_lts.x | bash -
RUN apt-get install -y nodejs

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -y libpng-dev libjpeg-dev curl libxi6 build-essential libgl1-mesa-glx
RUN curl -sL https://deb.nodesource.com/setup_lts.x | bash -
RUN apt-get install -y nodejs

WORKDIR /src
COPY ["Konektor/Konektor.csproj", "Konektor/"]
COPY ["Application/Application.csproj", "Application/"]
COPY ["Domain/Domain.csproj", "Domain/"]
COPY ["Infrastracture/Infrastracture.csproj", "Infrastracture/"]
RUN dotnet restore --no-cache "Konektor/Konektor.csproj"
COPY . .
WORKDIR "/src/Konektor"
RUN dotnet build "Konektor.csproj" -c Release -o /app/build

COPY ["ClientApp/package.json", "ClientApp/"]
COPY ["ClientApp/package-lock.json", "ClientApp/"]
WORKDIR /src/ClientApp

RUN npm install

FROM build AS publish
RUN dotnet publish "Konektor.csproj" -c Release -o /app/publish 

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Konektor.dll"]