-- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
--
-- Host: localhost    Database: medassist_db
-- ------------------------------------------------------
-- Server version	8.0.29-0ubuntu0.20.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Alerte`
--

DROP TABLE IF EXISTS `Alerte`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Alerte` (
  `IDAlerta` int NOT NULL AUTO_INCREMENT,
  `Gravitate` int DEFAULT NULL,
  `Mesaj` text,
  `IDSender` int DEFAULT NULL,
  PRIMARY KEY (`IDAlerta`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Alerte`
--

LOCK TABLES `Alerte` WRITE;
/*!40000 ALTER TABLE `Alerte` DISABLE KEYS */;
INSERT INTO `Alerte` VALUES (1,2,'E pe duca',1),(2,2,'E mai rau ca primul',2);
/*!40000 ALTER TABLE `Alerte` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DateMobile`
--

DROP TABLE IF EXISTS `DateMobile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DateMobile` (
  `IDDate` int NOT NULL AUTO_INCREMENT,
  `Data` datetime DEFAULT NULL,
  `Greutate` float DEFAULT NULL,
  `Glicemie` float DEFAULT NULL,
  `Tensiune_Mica` int DEFAULT NULL,
  `Tensiune_Mare` int DEFAULT NULL,
  `Temperatura` int DEFAULT NULL,
  `IDPacient` int DEFAULT NULL,
  PRIMARY KEY (`IDDate`),
  KEY `FK_Pacient` (`IDPacient`),
  CONSTRAINT `FK_Pacient` FOREIGN KEY (`IDPacient`) REFERENCES `Pacienti` (`IDPacient`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DateMobile`
--

LOCK TABLES `DateMobile` WRITE;
/*!40000 ALTER TABLE `DateMobile` DISABLE KEYS */;
INSERT INTO `DateMobile` VALUES (0,'1000-01-01 00:00:00',60,50,100,100,37,1),(1,'1000-01-01 00:00:00',60,50,100,100,37,1);
/*!40000 ALTER TABLE `DateMobile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DateSenzori`
--

DROP TABLE IF EXISTS `DateSenzori`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DateSenzori` (
  `IDSenzor` int NOT NULL AUTO_INCREMENT,
  `Puls` int DEFAULT NULL,
  `Lumina` int DEFAULT NULL,
  `Alerta_Proximitate` tinyint DEFAULT NULL,
  `Temperatura_Amb` float DEFAULT NULL,
  `Hum_Alert` tinyint DEFAULT NULL,
  `Gas_Alert` tinyint DEFAULT NULL,
  `IDPacient` int DEFAULT NULL,
  PRIMARY KEY (`IDSenzor`),
  KEY `FK_Pacient2` (`IDPacient`),
  CONSTRAINT `FK_Pacient2` FOREIGN KEY (`IDPacient`) REFERENCES `Pacienti` (`IDPacient`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DateSenzori`
--

LOCK TABLES `DateSenzori` WRITE;
/*!40000 ALTER TABLE `DateSenzori` DISABLE KEYS */;
INSERT INTO `DateSenzori` VALUES (0,100,50,1,28.5,0,0,1),(1,69,50,1,25.5,0,0,1),(2,80,60,1,25,0,0,2),(3,70,50,1,31,0,1,2);
/*!40000 ALTER TABLE `DateSenzori` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `FisaMed`
--

DROP TABLE IF EXISTS `FisaMed`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `FisaMed` (
  `IDFisa` int NOT NULL AUTO_INCREMENT,
  `Istoric_Medical` text,
  `Lista_Alergii` text,
  `Recomandari` text,
  `Schema_Medicatie` text,
  PRIMARY KEY (`IDFisa`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `FisaMed`
--

LOCK TABLES `FisaMed` WRITE;
/*!40000 ALTER TABLE `FisaMed` DISABLE KEYS */;
INSERT INTO `FisaMed` VALUES (1,'a1','a2','a3','a4'),(2,'b1','b2','b3','b4'),(3,'c1','c2','c3','c4'),(4,'d1','d2','d3','d4');
/*!40000 ALTER TABLE `FisaMed` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Ingrijitori`
--

DROP TABLE IF EXISTS `Ingrijitori`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Ingrijitori` (
  `IDIngrijitor` int NOT NULL AUTO_INCREMENT,
  `Nume` text,
  `Prenume` text,
  `Username` text,
  PRIMARY KEY (`IDIngrijitor`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Ingrijitori`
--

LOCK TABLES `Ingrijitori` WRITE;
/*!40000 ALTER TABLE `Ingrijitori` DISABLE KEYS */;
INSERT INTO `Ingrijitori` VALUES (1,'Marinescu','Mihai','Yes'),(2,'Padurescu','Alexandra','Yes');
/*!40000 ALTER TABLE `Ingrijitori` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Medici`
--

DROP TABLE IF EXISTS `Medici`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Medici` (
  `IDMedic` int NOT NULL AUTO_INCREMENT,
  `Nume` text,
  `Prenume` text,
  `Username` text,
  PRIMARY KEY (`IDMedic`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Medici`
--

LOCK TABLES `Medici` WRITE;
/*!40000 ALTER TABLE `Medici` DISABLE KEYS */;
INSERT INTO `Medici` VALUES (1,'Gheorghita','Vasilescul','Gheorghita'),(2,'Popescu','Ana','Popescu');
/*!40000 ALTER TABLE `Medici` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Pacienti`
--

DROP TABLE IF EXISTS `Pacienti`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Pacienti` (
  `IDPacient` int NOT NULL AUTO_INCREMENT,
  `Nume` text,
  `Prenume` text,
  `Varsta` int DEFAULT NULL,
  `Adresa` text,
  `NrTel` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL,
  `Mail` text,
  `Profesia` text,
  `LocDeMunca` text,
  `Username` text,
  `IDMedic` int DEFAULT NULL,
  `IDIngrijitor` int DEFAULT NULL,
  `IDSupraveghetor` int DEFAULT NULL,
  `IDFisa` int DEFAULT NULL,
  PRIMARY KEY (`IDPacient`),
  KEY `FK_Medic` (`IDMedic`),
  KEY `FK_Ingrijitor` (`IDIngrijitor`),
  KEY `FK_Supraveghetor` (`IDSupraveghetor`),
  KEY `IDFisa` (`IDFisa`),
  CONSTRAINT `FK_Ingrijitor` FOREIGN KEY (`IDIngrijitor`) REFERENCES `Ingrijitori` (`IDIngrijitor`),
  CONSTRAINT `FK_Medic` FOREIGN KEY (`IDMedic`) REFERENCES `Medici` (`IDMedic`),
  CONSTRAINT `FK_Supraveghetor` FOREIGN KEY (`IDSupraveghetor`) REFERENCES `Supraveghetori` (`IDSupraveghetor`),
  CONSTRAINT `Pacienti_ibfk_1` FOREIGN KEY (`IDFisa`) REFERENCES `FisaMed` (`IDFisa`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Pacienti`
--

LOCK TABLES `Pacienti` WRITE;
/*!40000 ALTER TABLE `Pacienti` DISABLE KEYS */;
INSERT INTO `Pacienti` VALUES (1,'Ionescu','Radu',32,'strada Linistei','722124571','radu@gmail.com','medic',' Premier','Yes',2,1,1,1),(2,'Alexandrescu','Maria',47,'timisoara','745663981','maria@gmail.com','vanzator','Lidl','Yes',2,2,2,2),(3,'Grigorescu','Matei',28,'giroc','799765643','matei@gmail.com','pilot','Tarom','Yes',1,1,1,3),(4,'Mihai','Andreea',50,'ghiroda','723447283','andreea@gmail.com','judecator','tribunal','Yes',1,2,2,4),(9,'Stanescu','Paul',30,'Timisoara','752359305','paul@gmail.com','Avocat','LawOne','paul_stan',1,1,1,1);
/*!40000 ALTER TABLE `Pacienti` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Senders`
--

DROP TABLE IF EXISTS `Senders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Senders` (
  `IDSender` int NOT NULL AUTO_INCREMENT,
  `Username` text,
  `Type` text,
  PRIMARY KEY (`IDSender`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Senders`
--

LOCK TABLES `Senders` WRITE;
/*!40000 ALTER TABLE `Senders` DISABLE KEYS */;
INSERT INTO `Senders` VALUES (1,'popescu','medic'),(3,'username1','ingrijitor');
/*!40000 ALTER TABLE `Senders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Supraveghetori`
--

DROP TABLE IF EXISTS `Supraveghetori`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Supraveghetori` (
  `IDSupraveghetor` int NOT NULL AUTO_INCREMENT,
  `Nume` text,
  `Prenume` text,
  `Username` text,
  PRIMARY KEY (`IDSupraveghetor`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Supraveghetori`
--

LOCK TABLES `Supraveghetori` WRITE;
/*!40000 ALTER TABLE `Supraveghetori` DISABLE KEYS */;
INSERT INTO `Supraveghetori` VALUES (1,'Rica','Monica','Yes'),(2,'Damian','Eric','Yes');
/*!40000 ALTER TABLE `Supraveghetori` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Users`
--

DROP TABLE IF EXISTS `Users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Users` (
  `IDUser` int NOT NULL AUTO_INCREMENT,
  `Username` text,
  `Password` text,
  PRIMARY KEY (`IDUser`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Users`
--

LOCK TABLES `Users` WRITE;
/*!40000 ALTER TABLE `Users` DISABLE KEYS */;
INSERT INTO `Users` VALUES (5,'Gheorghita','Gheorghita'),(6,'Popescu','Popescu');
/*!40000 ALTER TABLE `Users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-25 12:14:57
