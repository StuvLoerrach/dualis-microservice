-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: localhost    Database: dualis
-- ------------------------------------------------------
-- Server version	5.5.5-10.5.9-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `organization_fk` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `course_FK` (`organization_fk`),
  CONSTRAINT `course_FK` FOREIGN KEY (`organization_fk`) REFERENCES `organization` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES (1,'TIF18A-G2',1),(2,'TIF18A-G1',1);
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `enrollment`
--

DROP TABLE IF EXISTS `enrollment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `enrollment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `grade` varchar(255) DEFAULT NULL,
  `status` varchar(255) NOT NULL,
  `module_fk` int(11) NOT NULL,
  `semester_fk` int(11) NOT NULL,
  `student_fk` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `enrollment_FK` (`module_fk`),
  KEY `enrollment_FK_1` (`semester_fk`),
  KEY `enrollment_FK_2` (`student_fk`),
  CONSTRAINT `enrollment_FK` FOREIGN KEY (`module_fk`) REFERENCES `module` (`id`),
  CONSTRAINT `enrollment_FK_1` FOREIGN KEY (`semester_fk`) REFERENCES `semester` (`id`),
  CONSTRAINT `enrollment_FK_2` FOREIGN KEY (`student_fk`) REFERENCES `student` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `enrollment`
--

LOCK TABLES `enrollment` WRITE;
/*!40000 ALTER TABLE `enrollment` DISABLE KEYS */;
INSERT INTO `enrollment` VALUES (1,'2,0','bestanden',1,1,4),(2,'1,8','bestanden',2,3,4),(3,'noch nicht gesetzt','',3,6,4),(4,'2,2','bestanden',3,7,4),(5,'2,0','bestanden',2,3,3),(6,'1,4','bestanden',2,3,5),(7,'4,3','nicht bestanden',2,3,6);
/*!40000 ALTER TABLE `enrollment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lecture`
--

DROP TABLE IF EXISTS `lecture`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lecture` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `no` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `weighting` varchar(255) NOT NULL,
  `exam_type` varchar(255) NOT NULL,
  `module_fk` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lecture_FK` (`module_fk`),
  CONSTRAINT `lecture_FK` FOREIGN KEY (`module_fk`) REFERENCES `module` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lecture`
--

LOCK TABLES `lecture` WRITE;
/*!40000 ALTER TABLE `lecture` DISABLE KEYS */;
INSERT INTO `lecture` VALUES (1,'T3INF2005.1','Rechnerarchitekturen 1','6/16','Klausurarbeit',1),(2,'T3INF2005.2','Betriebssysteme','6/16','Klausurarbeit',1),(3,'T3INF2005.3','Systemnahe Programmierung 1','4/16','Projekt',1),(4,'T3INF4316.1','Agile Prozessmodelle','50%','Schriftl. Ausarbeitung',2),(5,'T3INF4392.1','Klassische Vorgehensmodelle','50%','Schriftl. Ausarbeitung',2),(6,'T3INF1001.1','Lineare Algebra','50%','Klausurarbeit',3),(7,'T3INF1001.2','Analysis','50%','Klausurarbeit',3);
/*!40000 ALTER TABLE `lecture` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lecture_result`
--

DROP TABLE IF EXISTS `lecture_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lecture_result` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `grade` varchar(255) DEFAULT NULL,
  `presence` tinyint(1) DEFAULT NULL,
  `enrollment_fk` int(11) NOT NULL,
  `lecture_fk` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `lecture_result_FK` (`enrollment_fk`),
  KEY `lecture_result_FK_1` (`lecture_fk`),
  CONSTRAINT `lecture_result_FK` FOREIGN KEY (`enrollment_fk`) REFERENCES `enrollment` (`id`),
  CONSTRAINT `lecture_result_FK_1` FOREIGN KEY (`lecture_fk`) REFERENCES `lecture` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lecture_result`
--

LOCK TABLES `lecture_result` WRITE;
/*!40000 ALTER TABLE `lecture_result` DISABLE KEYS */;
INSERT INTO `lecture_result` VALUES (1,'85,2',1,1,1),(2,'75,0',1,1,2),(3,'71,4',1,1,3),(4,'1,7',1,2,4),(5,'1,9',1,2,5),(6,'2,0',1,3,6),(7,'2,4',1,4,7),(8,'1,9',1,5,4),(9,'2,1',1,5,5),(10,'1,2',1,6,4),(11,'1,6',1,6,5),(12,'4,0',1,7,4),(13,'4,6',0,7,5);
/*!40000 ALTER TABLE `lecture_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `module`
--

DROP TABLE IF EXISTS `module`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `module` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `no` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `credits` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `module`
--

LOCK TABLES `module` WRITE;
/*!40000 ALTER TABLE `module` DISABLE KEYS */;
INSERT INTO `module` VALUES (1,'T3INF2005','Technische Informatik II','8,0'),(2,'T3INF4392','Vorgehensmodelle','5,0'),(3,'T3INF1001','Mathematik I','8,0');
/*!40000 ALTER TABLE `module` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organization`
--

DROP TABLE IF EXISTS `organization`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `organization` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organization`
--

LOCK TABLES `organization` WRITE;
/*!40000 ALTER TABLE `organization` DISABLE KEYS */;
INSERT INTO `organization` VALUES (1,'DHBW Lörrach');
/*!40000 ALTER TABLE `organization` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `semester`
--

DROP TABLE IF EXISTS `semester`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `semester` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_wintersemester` tinyint(1) NOT NULL,
  `year` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `semester`
--

LOCK TABLES `semester` WRITE;
/*!40000 ALTER TABLE `semester` DISABLE KEYS */;
INSERT INTO `semester` VALUES (1,1,'2019'),(2,0,'2020'),(3,1,'2020'),(4,0,'2021'),(5,1,'2021'),(6,1,'2018'),(7,0,'2019');
/*!40000 ALTER TABLE `semester` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `course_fk` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `student_FK` (`course_fk`),
  CONSTRAINT `student_FK` FOREIGN KEY (`course_fk`) REFERENCES `course` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--

LOCK TABLES `student` WRITE;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
INSERT INTO `student` VALUES (3,'scuderir@dhbw-loerrach.de','euifhff',1),(4,'kaiseand@dhbw-loerrach.de','gopgj',1),(5,'behrends@dhbw-loerrach.de','abcd',2),(6,'peter@dhbw-loerrach.de','1234',1);
/*!40000 ALTER TABLE `student` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'dualis'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-05-07 14:51:07
