USE [Widgets]
GO

/****** Object:  Table [dbo].[Widget]    Script Date: 16/10/2017 10:54:59 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[Widget](
	[Id] [int] IDENTITY(1,1) NOT NULL,
	[Name] [varchar](150) NOT NULL,
	[Color] [varchar](50) NOT NULL,
	[Price] [decimal](18, 2) NOT NULL,
	[Inventory] [int] NOT NULL,
	[Melts] [bit] NOT NULL,
 CONSTRAINT [PK__Widget__3214EC071228009B] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO


