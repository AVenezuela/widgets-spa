USE [Widget]
GO

/****** Object:  Table [dbo].[Widgets]    Script Date: 12/10/2017 02:59:12 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[Widgets](
	[Id] [int] IDENTITY(1,1) NOT NULL,
	[Color] [varchar](50) NOT NULL,
	[Price] [decimal](18, 2) NOT NULL,
	[Inventory] [int] NOT NULL,
	[Melts] [bit] NOT NULL,
 CONSTRAINT [PK_Widgets] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO


