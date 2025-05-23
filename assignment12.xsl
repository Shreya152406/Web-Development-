<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:output method="html"/>
<xsl:template match="/">
<html> 
<body>
  <h2>Chennai Super Kings</h2>
  <table border="1">
    <tr bgcolor="#9acd32">
      <th style="text-align:left">Player</th>
      <th style="text-align:left">Role</th>
    </tr>
    <xsl:for-each select="catalog/cd">
    <tr>
      <td><xsl:value-of select="player"/></td>
      <td><xsl:value-of select="role"/></td>
    </tr>
    </xsl:for-each>
  </table>
</body>
</html>
</xsl:template>
</xsl:stylesheet>