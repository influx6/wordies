package abbrvs

//go:generate go run generate.go

var (
	abbreviations = map[string][]string{
		
		"A.D.": []string{ "Anno Domini", },
		
		"A.V.": []string{ "Authorized Version", },
		
		"Abbrev.": []string{ "abbreviation(s)", },
		
		"Abd.": []string{ "Aberdeen", },
		
		"Aberd.": []string{ "Aberdeen", },
		
		"Aberdeensh.": []string{ "Aberdeenshire", },
		
		"Abol.": []string{ "abolition", },
		
		"Aborig.": []string{ "aboriginal", },
		
		"Abp.": []string{ "archbishop", },
		
		"Abr.": []string{ "abridged", },
		
		"Abridg.": []string{ "abridged", },
		
		"Abridgem.": []string{ "abridgement", },
		
		"Absol.": []string{ "absolute", },
		
		"Abst.": []string{ "abstract(s)", },
		
		"Abstr.": []string{ "abstract(s)", },
		
		"Acad.": []string{ "academia"," academy"," academic(al)", },
		
		"Acc.": []string{ "account(s)", },
		
		"Accept.": []string{ "acceptance", },
		
		"Accomm.": []string{ "accommodation", },
		
		"Accompl.": []string{ "accomplished", },
		
		"Accs.": []string{ "accounts", },
		
		"Acct.": []string{ "account", },
		
		"Accts.": []string{ "accounts", },
		
		"Achievem.": []string{ "achievement", },
		
		"Add.": []string{ "addendum"," addenda", },
		
		"Addit.": []string{ "addition(s)", },
		
		"Addr.": []string{ "address", },
		
		"Adm.": []string{ "admiralty", },
		
		"Admin.": []string{ "administration"," administrative", },
		
		"Admir.": []string{ "admiralty", },
		
		"Admon.": []string{ "admonition", },
		
		"Admonit.": []string{ "admonition", },
		
		"Adv.": []string{ "advocate", },
		
		"Advancem.": []string{ "advancement(s)", },
		
		"Advert.": []string{ "advertisement", },
		
		"Advoc.": []string{ "advocate", },
		
		"Advt.": []string{ "advertisement", },
		
		"Advts.": []string{ "advertisements", },
		
		"Aerodynam.": []string{ "aerodynamics", },
		
		"Aeronaut.": []string{ "aeronautics"," aeronautical", },
		
		"Aff.": []string{ "affairs", },
		
		"Affect.": []string{ "affection(s)", },
		
		"Afr.": []string{ "Africa(n)", },
		
		"Agric.": []string{ "agriculture"," agricultural", },
		
		"Alch.": []string{ "alchemy", },
		
		"Alg.": []string{ "algebra", },
		
		"Alleg.": []string{ "allegiance", },
		
		"Allit.": []string{ "alliterative", },
		
		"Alm.": []string{ "almanac", },
		
		"Alph.": []string{ "alphabet"," alphabetical", },
		
		"Amer.": []string{ "America"," American", },
		
		"Anal.": []string{ "analysis", },
		
		"Analyt.": []string{ "analytic(al)", },
		
		"Anat.": []string{ "anatomy"," anatomical", },
		
		"Anc.": []string{ "ancient", },
		
		"Anecd.": []string{ "anecdotes", },
		
		"Ang.": []string{ "Anglian", },
		
		"Angl.": []string{ "Anglian", },
		
		"Anglo-Ind.": []string{ "Anglo-Indian", },
		
		"Anim.": []string{ "animal(s)", },
		
		"Ann.": []string{ "annual", },
		
		"Anniv.": []string{ "anniversary", },
		
		"Annot.": []string{ "annotation", },
		
		"Anon.": []string{ "(in author names) Anonymous", },
		
		"Answ.": []string{ "answer", },
		
		"Ant.": []string{ "antiquities", },
		
		"Anthrop.": []string{ "anthropology", },
		
		"Anthropol.": []string{ "anthropology"," anthropological", },
		
		"Antiq.": []string{ "antiquity"," antiquities", },
		
		"Apoc.": []string{ "apocalyptic(al)", },
		
		"Apol.": []string{ "apology"," apologies", },
		
		"App.": []string{ "appendix", },
		
		"Appl.": []string{ "applied", },
		
		"Applic.": []string{ "application(s)", },
		
		"Apr.": []string{ "April", },
		
		"Arab.": []string{ "Arabic", },
		
		"Arb.": []string{ "(in author names) Arber", },
		
		"Arch.": []string{ "archive(s)", },
		
		"Archaeol.": []string{ "archaeology"," archaeological", },
		
		"Archipel.": []string{ "archipelago", },
		
		"Archit.": []string{ "architecture"," architectural", },
		
		"Argt.": []string{ "argument", },
		
		"Arith.": []string{ "arithmetic", },
		
		"Arithm.": []string{ "arithmetic"," arithmetical", },
		
		"Arrangem.": []string{ "arrangement", },
		
		"Artic.": []string{ "articulation", },
		
		"Artific.": []string{ "artificial", },
		
		"Artill.": []string{ "artillery", },
		
		"Ashm.": []string{ "(in author names) Ashmole", },
		
		"Assemb.": []string{ "assembly"," assemblies", },
		
		"Assoc.": []string{ "association", },
		
		"Assoc. Football": []string{ "Association Football", },
		
		"Assyriol.": []string{ "Assyriology", },
		
		"Astr.": []string{ "astronomy", },
		
		"Astrol.": []string{ "astrology", },
		
		"Astron.": []string{ "astronomy"," astronomic(al)", },
		
		"Astronaut.": []string{ "astronautics"," astronautical", },
		
		"Att.": []string{ "Attic", },
		
		"Attrib.": []string{ "attribute(s)", },
		
		"Aug.": []string{ "August", },
		
		"Austral.": []string{ "Australia(n)", },
		
		"Auth.": []string{ "author(s)", },
		
		"Autobiog.": []string{ "autobiography", },
		
		"Autobiogr.": []string{ "autobiography"," autobiographic(al)", },
		
		"Ayrsh.": []string{ "Ayrshire", },
		
		"B.C.": []string{ "British Columbia", },
		
		"BNC": []string{ "British National Corpus", },
		
		"Bacteriol.": []string{ "bacteriology"," bacteriologic(al)", },
		
		"Bedford.": []string{ "Bedfordshire", },
		
		"Bedfordsh.": []string{ "Bedfordshire", },
		
		"Bel & Dr.": []string{ "(in the Apocrypha) Bel and the Dragon", },
		
		"Belg.": []string{ "Belgian", },
		
		"Berks.": []string{ "Berkshire", },
		
		"Berksh.": []string{ "Berkshire", },
		
		"Berw.": []string{ "Berwickshire", },
		
		"Berwicksh.": []string{ "Berwickshire", },
		
		"Bibliogr.": []string{ "bibliography", },
		
		"Biochem.": []string{ "biochemistry"," biochemical", },
		
		"Biog.": []string{ "biography", },
		
		"Biogr.": []string{ "biography"," biographic(al)", },
		
		"Biol.": []string{ "biology"," biological", },
		
		"Bk.": []string{ "book", },
		
		"Bks.": []string{ "books", },
		
		"Bord.": []string{ "border", },
		
		"Bot.": []string{ "botany"," botanic(al)", },
		
		"Bp.": []string{ "Bishop", },
		
		"Braz.": []string{ "Brazilian", },
		
		"Brit.": []string{ "Britain"," British", },
		
		"Bucks.": []string{ "Buckinghamshire", },
		
		"Build.": []string{ "building", },
		
		"Bull.": []string{ "bulletin", },
		
		"Bur.": []string{ "bureau", },
		
		"C.": []string{ "county"," counties", },
		
		"Cal.": []string{ "calendar", },
		
		"Calc.": []string{ "calculus", },
		
		"Calend.": []string{ "calendar", },
		
		"Calif.": []string{ "California", },
		
		"Calligr.": []string{ "calligraphy", },
		
		"Camb.": []string{ "Cambridge", },
		
		"Cambr.": []string{ "Cambridge", },
		
		"Campanol.": []string{ "campanology", },
		
		"Canad.": []string{ "Canadian", },
		
		"Canterb.": []string{ "Canterbury", },
		
		"Capt.": []string{ "captain", },
		
		"Cartogr.": []string{ "cartography"," cartographic(al)", },
		
		"Catal.": []string{ "catalogue", },
		
		"Catech.": []string{ "catechism", },
		
		"Cath.": []string{ "Catholic", },
		
		"Cent.": []string{ "century", },
		
		"Ceram.": []string{ "ceramics", },
		
		"Cert.": []string{ "certificate", },
		
		"Certif.": []string{ "certificate", },
		
		"Ch.": []string{ "Church", },
		
		"Ch. Hist.": []string{ "Church history", },
		
		"Chamb.": []string{ "chamber", },
		
		"Char.": []string{ "character", },
		
		"Charac.": []string{ "character", },
		
		"Chas.": []string{ "Charles", },
		
		"Chem.": []string{ "chemistry"," chemical", },
		
		"Chem. Engin.": []string{ "chemical engineering", },
		
		"Chesh.": []string{ "Cheshire", },
		
		"Chr.": []string{ "(in the Bible) Chronicles", },
		
		"Chron.": []string{ "(in the Bible) Chronicles", },
		
		"Chronol.": []string{ "chronology", },
		
		"Chrons.": []string{ "chronicles", },
		
		"Cinematogr.": []string{ "cinematography"," cinematographic(al)", },
		
		"Circ.": []string{ "circle", },
		
		"Civ. Law": []string{ "civil law", },
		
		"Civil Engin.": []string{ "civil engineering", },
		
		"Cl.": []string{ "club", },
		
		"Class.": []string{ "classical", },
		
		"Class. Antiq.": []string{ "classical antiquities", },
		
		"Classif.": []string{ "classification", },
		
		"Climatol.": []string{ "climatology", },
		
		"Clin.": []string{ "clinical", },
		
		"Col.": []string{ "(in the Bible) Colossians", },
		
		"Coll.": []string{ "college", },
		
		"Collect.": []string{ "collection", },
		
		"Colloq.": []string{ "colloquies", },
		
		"Coloss.": []string{ "(in the Bible) Colossians", },
		
		"Com.": []string{ "commission", },
		
		"Comb.": []string{ "combination(s)", },
		
		"Combs.": []string{ "combinations", },
		
		"Comm.": []string{ "committee", },
		
		"Comm. Law": []string{ "commercial law", },
		
		"Commandm.": []string{ "commandment", },
		
		"Commend.": []string{ "commendation", },
		
		"Commerc.": []string{ "commercial", },
		
		"Commiss.": []string{ "commission", },
		
		"Commonw.": []string{ "commonwealth", },
		
		"Communic.": []string{ "communication(s)", },
		
		"Comp.": []string{ "company", },
		
		"Comp. Anat.": []string{ "comparative anatomy", },
		
		"Compan.": []string{ "companion(s)", },
		
		"Compar.": []string{ "comparative", },
		
		"Compend.": []string{ "compendium"," compendious", },
		
		"Compl.": []string{ "complete", },
		
		"Compos.": []string{ "composition", },
		
		"Conc.": []string{ "concise", },
		
		"Conch.": []string{ "conchology", },
		
		"Concl.": []string{ "conclusion", },
		
		"Conf.": []string{ "confession", },
		
		"Confid.": []string{ "confidential", },
		
		"Confl.": []string{ "conflict", },
		
		"Confut.": []string{ "confutation", },
		
		"Congr.": []string{ "congress", },
		
		"Congreg.": []string{ "congregation", },
		
		"Congress.": []string{ "congressional", },
		
		"Conn.": []string{ "Connecticut", },
		
		"Consc.": []string{ "conscience", },
		
		"Consecr.": []string{ "consecrated"," consecration", },
		
		"Consid.": []string{ "consideration(s)", },
		
		"Consol.": []string{ "consolation", },
		
		"Constit.": []string{ "constitution(al)", },
		
		"Constit. Hist.": []string{ "constitutional history", },
		
		"Constr.": []string{ "construction", },
		
		"Contemp.": []string{ "contemporary", },
		
		"Contempl.": []string{ "contemplation", },
		
		"Contend.": []string{ "contending(s)", },
		
		"Content.": []string{ "contention", },
		
		"Contin.": []string{ "continuation", },
		
		"Contradict.": []string{ "contradiction", },
		
		"Contrib.": []string{ "contribution", },
		
		"Controv.": []string{ "controversy", },
		
		"Conv.": []string{ "convention", },
		
		"Convent.": []string{ "convention", },
		
		"Conversat.": []string{ "conversation", },
		
		"Convoc.": []string{ "convocation", },
		
		"Cor.": []string{ "(in the Bible) Corinthians", },
		
		"Cornw.": []string{ "Cornwall", },
		
		"Coron.": []string{ "coronation", },
		
		"Corr.": []string{ "correspondence", },
		
		"Corresp.": []string{ "correspondence", },
		
		"Counc.": []string{ "council", },
		
		"Courtsh.": []string{ "courtship", },
		
		"Craniol.": []string{ "craniology", },
		
		"Craniom.": []string{ "craniometry", },
		
		"Crim.": []string{ "criminal", },
		
		"Crim. Law": []string{ "criminal law", },
		
		"Crit.": []string{ "criticism", },
		
		"Crt.": []string{ "court", },
		
		"Crts.": []string{ "courts", },
		
		"Cryptogr.": []string{ "cryptography", },
		
		"Crystallogr.": []string{ "crystallography", },
		
		"Ct.": []string{ "court", },
		
		"Cumb.": []string{ "Cumberland", },
		
		"Cumberld.": []string{ "Cumberland", },
		
		"Cumbld.": []string{ "Cumberland", },
		
		"Cycl.": []string{ "cyclopaedia", },
		
		"Cytol.": []string{ "cytology"," cytological", },
		
		"D.C.": []string{ "District of Columbia", },
		
		"Dan.": []string{ "(in the Bible) Daniel", },
		
		"Dau.": []string{ "daughter", },
		
		"Deb.": []string{ "debate", },
		
		"Dec.": []string{ "December", },
		
		"Declar.": []string{ "declaration", },
		
		"Ded.": []string{ "dedication", },
		
		"Def.": []string{ "defence", },
		
		"Deliv.": []string{ "deliverance", },
		
		"Demonstr.": []string{ "demonstration", },
		
		"Dep.": []string{ "deputy", },
		
		"Depred.": []string{ "depredation", },
		
		"Depredat.": []string{ "depredation", },
		
		"Dept.": []string{ "department", },
		
		"Derbysh.": []string{ "Derbyshire", },
		
		"Descr.": []string{ "description"," descriptive", },
		
		"Deut.": []string{ "(in the Bible) Deuteronomy", },
		
		"Devel.": []string{ "development", },
		
		"Devonsh.": []string{ "Devonshire", },
		
		"Dial.": []string{ "dialect", },
		
		"Dict.": []string{ "dictionary", },
		
		"Diffic.": []string{ "difficult"," difficulty", },
		
		"Direct.": []string{ "direction", },
		
		"Dis.": []string{ "disease(s)", },
		
		"Disc.": []string{ "discourse", },
		
		"Discipl.": []string{ "discipline", },
		
		"Discov.": []string{ "discovery", },
		
		"Discrim.": []string{ "discrimination", },
		
		"Discuss.": []string{ "discussion", },
		
		"Diss.": []string{ "dissertation", },
		
		"Dist.": []string{ "district", },
		
		"Distemp.": []string{ "distemper", },
		
		"Distill.": []string{ "distilling", },
		
		"Distrib.": []string{ "distribution", },
		
		"Div.": []string{ "division", },
		
		"Divers.": []string{ "diversity", },
		
		"Dk.": []string{ "Duke", },
		
		"Doc.": []string{ "document(s)", },
		
		"Doctr.": []string{ "doctrine", },
		
		"Domest.": []string{ "domestic", },
		
		"Dr.": []string{ "Doctor", },
		
		"Durh.": []string{ "Durham", },
		
		"E.": []string{ "east"," eastern", },
		
		"E. Afr.": []string{ "East Africa(n)", },
		
		"E. Angl.": []string{ "East Anglia(n)", },
		
		"E. Anglian": []string{ "East Anglian", },
		
		"E. Ind.": []string{ "East Indies"," East Indian", },
		
		"E.D.D.": []string{ "English Dialect Dictionary", },
		
		"E.E.T.S.": []string{ "Early English Text Society", },
		
		"East Ind.": []string{ "East Indies"," East Indian", },
		
		"Eccl.": []string{ "(in the Bible) Ecclesiastes", },
		
		"Eccl. Hist.": []string{ "ecclesiastical history", },
		
		"Eccl. Law": []string{ "ecclesiastical law", },
		
		"Eccles.": []string{ "(in the Bible) Ecclesiastes", },
		
		"Ecclus.": []string{ "(in the Apocrypha) Ecclesiasticus", },
		
		"Ecol.": []string{ "ecology"," ecological", },
		
		"Econ.": []string{ "economics"," economic(al)", },
		
		"Ed.": []string{ "Edward", },
		
		"Edin.": []string{ "Edinburgh", },
		
		"Edinb.": []string{ "Edinburgh", },
		
		"Educ.": []string{ "education"," education(al)", },
		
		"Edw.": []string{ "Edward", },
		
		"Egypt.": []string{ "Egyptian", },
		
		"Egyptol.": []string{ "Egyptology", },
		
		"Electr.": []string{ "electricity"," electric(al)", },
		
		"Electr. Engin.": []string{ "electrical engineering", },
		
		"Electro-magn.": []string{ "electro-magnetism", },
		
		"Electro-physiol.": []string{ "electro-physiology", },
		
		"Elem.": []string{ "element(s)"," elementary", },
		
		"Eliz.": []string{ "Elizabeth", },
		
		"Elizab.": []string{ "Elizabethan", },
		
		"Emb.": []string{ "embassy", },
		
		"Embryol.": []string{ "embryology"," embryologic(al)", },
		
		"Encycl.": []string{ "encyclopaedia"," encyclopaedic", },
		
		"Encycl. Brit.": []string{ "Encyclopaedia Britannica", },
		
		"Encycl. Metrop.": []string{ "Encyclopaedia Metropolitana", },
		
		"Eng.": []string{ "England"," English", },
		
		"Engin.": []string{ "engineering", },
		
		"Englishw.": []string{ "Englishwomen", },
		
		"Enq.": []string{ "enquiry", },
		
		"Ent.": []string{ "entomology"," entomological", },
		
		"Enthus.": []string{ "enthusiasm", },
		
		"Entom.": []string{ "entomology"," entomological", },
		
		"Entomol.": []string{ "entomology"," entomological", },
		
		"Enzymol.": []string{ "enzymology", },
		
		"Ep.": []string{ "epistle", },
		
		"Eph.": []string{ "(in the Bible) Ephesians", },
		
		"Ephes.": []string{ "(in the Bible) Ephesians", },
		
		"Epil.": []string{ "epilogue", },
		
		"Episc.": []string{ "episcopacy", },
		
		"Epist.": []string{ "epistle", },
		
		"Epit.": []string{ "epitaph", },
		
		"Equip.": []string{ "equipment", },
		
		"Esd.": []string{ "(in the Apocrypha) Esdras", },
		
		"Ess.": []string{ "essay(s)", },
		
		"Essent.": []string{ "essential", },
		
		"Establ.": []string{ "establishment", },
		
		"Esth.": []string{ "(in the Bible) Esther", },
		
		"Ethnol.": []string{ "ethnology", },
		
		"Etymol.": []string{ "etymology"," etymological", },
		
		"Eval.": []string{ "evaluation", },
		
		"Evang.": []string{ "evangelical", },
		
		"Even.": []string{ "evening", },
		
		"Evid.": []string{ "evidence", },
		
		"Evol.": []string{ "evolution"," evolutionary", },
		
		"Ex. doc.": []string{ "executive document", },
		
		"Exalt.": []string{ "exaltation", },
		
		"Exam.": []string{ "examination", },
		
		"Exch.": []string{ "exchequer", },
		
		"Exec.": []string{ "executive", },
		
		"Exerc.": []string{ "exercise(s)", },
		
		"Exhib.": []string{ "exhibition"," exhibit(s)", },
		
		"Exod.": []string{ "(in the Bible) Exodus", },
		
		"Exped.": []string{ "expedition", },
		
		"Exper.": []string{ "experiment"," experimental", },
		
		"Explan.": []string{ "explanation", },
		
		"Explic.": []string{ "explication", },
		
		"Explor.": []string{ "exploration", },
		
		"Expos.": []string{ "exposition"," expository", },
		
		"Ezek.": []string{ "(in the Bible) Ezekiel", },
		
		"Fab.": []string{ "fable(s)", },
		
		"Fam.": []string{ "family", },
		
		"Farew.": []string{ "farewell", },
		
		"Feb.": []string{ "February", },
		
		"Ff.": []string{ "folios", },
		
		"Fifesh.": []string{ "Fifeshire", },
		
		"Footpr.": []string{ "footprint", },
		
		"Forfarsh.": []string{ "Forfarshire", },
		
		"Fortif.": []string{ "fortification", },
		
		"Fortn.": []string{ "fortnightly", },
		
		"Found.": []string{ "foundation(s)", },
		
		"Fr.": []string{ "from", },
		
		"Fragm.": []string{ "fragment(s)", },
		
		"Fratern.": []string{ "fraternity", },
		
		"Friendsh.": []string{ "friendship", },
		
		"Fund.": []string{ "fundamental(s)", },
		
		"Furnit.": []string{ "furniture", },
		
		"Gal.": []string{ "(in the Bible) Galatians", },
		
		"Gard.": []string{ "garden"," gardening", },
		
		"Gastron.": []string{ "gastronomy", },
		
		"Gaz.": []string{ "gazette", },
		
		"Gd.": []string{ "good", },
		
		"Gen.": []string{ "(in the Bible) Genesis", },
		
		"Geo.": []string{ "George", },
		
		"Geog.": []string{ "geography", },
		
		"Geogr.": []string{ "geography"," geographic(al)", },
		
		"Geol.": []string{ "geology"," geologic(al)", },
		
		"Geom.": []string{ "geometry"," geometric(al)", },
		
		"Geomorphol.": []string{ "geomorphology", },
		
		"Ger.": []string{ "German", },
		
		"Glac.": []string{ "glacier(s)", },
		
		"Glasg.": []string{ "Glasgow", },
		
		"Glos.": []string{ "Gloucestershire", },
		
		"Gloss.": []string{ "glossary", },
		
		"Glouc.": []string{ "Gloucestershire", },
		
		"Gloucestersh.": []string{ "Gloucestershire", },
		
		"Gosp.": []string{ "gospel(s)", },
		
		"Gov.": []string{ "government", },
		
		"Govt.": []string{ "government", },
		
		"Gr.": []string{ "grammar", },
		
		"Gram.": []string{ "grammar", },
		
		"Gramm. Analysis": []string{ "grammatical analysis", },
		
		"Gt.": []string{ "great", },
		
		"Gynaecol.": []string{ "gynaecology", },
		
		"Hab.": []string{ "(in the Bible) Habakkuk", },
		
		"Haematol.": []string{ "haematology", },
		
		"Hag.": []string{ "(in the Bible) Haggai", },
		
		"Hampsh.": []string{ "Hampshire", },
		
		"Handbk.": []string{ "handbook", },
		
		"Hants.": []string{ "Hampshire", },
		
		"Heb.": []string{ "(in the Bible) Hebrews", },
		
		"Hebr.": []string{ "(in the Bible) Hebrews", },
		
		"Hen.": []string{ "Henry", },
		
		"Her.": []string{ "heraldry", },
		
		"Herb.": []string{ "herbalism", },
		
		"Heref.": []string{ "Herefordshire", },
		
		"Hereford.": []string{ "Herefordshire", },
		
		"Herefordsh.": []string{ "Herefordshire", },
		
		"Hertfordsh.": []string{ "Hertfordshire", },
		
		"Hierogl.": []string{ "Hieroglyphic", },
		
		"Hist.": []string{ "history"," historic(al)", },
		
		"Histol.": []string{ "histology"," histologic(al)", },
		
		"Hom.": []string{ "homilies", },
		
		"Horol.": []string{ "horology", },
		
		"Hort.": []string{ "horticulture"," horticultural", },
		
		"Hos.": []string{ "(in the Bible) Hosea", },
		
		"Hosp.": []string{ "hospital", },
		
		"Househ.": []string{ "household", },
		
		"Housek.": []string{ "housekeeping", },
		
		"Husb.": []string{ "husbandry", },
		
		"Hydraul.": []string{ "hydraulic", },
		
		"Hydrol.": []string{ "hydrology"," hydrological", },
		
		"Ichth.": []string{ "ichthyology", },
		
		"Icthyol.": []string{ "icthyology", },
		
		"Ideol.": []string{ "ideology"," ideological", },
		
		"Idol.": []string{ "idolatry", },
		
		"Illustr.": []string{ "illustrated"," illustration(s)", },
		
		"Imag.": []string{ "imaginary"," imagination", },
		
		"Immunol.": []string{ "immunology", },
		
		"Impr.": []string{ "improved", },
		
		"Inaug.": []string{ "inaugural", },
		
		"Inclos.": []string{ "inclosure", },
		
		"Ind.": []string{ "industry", },
		
		"Industr.": []string{ "industrial", },
		
		"Industr. Rel.": []string{ "industrial relations", },
		
		"Infl.": []string{ "influence", },
		
		"Innoc.": []string{ "innocence"," innocent", },
		
		"Inorg.": []string{ "inorganic", },
		
		"Inq.": []string{ "inquiry", },
		
		"Inst.": []string{ "institute"," institution", },
		
		"Instr.": []string{ "instruction", },
		
		"Intell.": []string{ "intelligence"," intelligent", },
		
		"Intellect.": []string{ "intellectual", },
		
		"Interc.": []string{ "intercourse", },
		
		"Interl.": []string{ "interlude", },
		
		"Internat.": []string{ "international", },
		
		"Interpr.": []string{ "interpretation", },
		
		"Intro.": []string{ "introduction", },
		
		"Introd.": []string{ "introduction"," introductory", },
		
		"Inv.": []string{ "inventory"," inventories", },
		
		"Invent.": []string{ "inventory"," inventories", },
		
		"Invert. Zool.": []string{ "invertebrate zoology", },
		
		"Invertebr.": []string{ "invertebrate"," invertebrated", },
		
		"Investig.": []string{ "investigation", },
		
		"Investm.": []string{ "investment", },
		
		"Invoc.": []string{ "invocation", },
		
		"Ir.": []string{ "Irish", },
		
		"Irel.": []string{ "Ireland", },
		
		"Isa.": []string{ "(in the Bible) Isaiah", },
		
		"Ital.": []string{ "Italian", },
		
		"Jahrb.": []string{ "Jahrbuch", },
		
		"Jam.": []string{ "(in author names) Jamieson", },
		
		"Jan.": []string{ "January", },
		
		"Jap.": []string{ "Japanese", },
		
		"Jas.": []string{ "(in the Bible) James", },
		
		"Jer.": []string{ "(in the Bible) Jeremiah", },
		
		"Josh.": []string{ "(in the Bible) Joshua", },
		
		"Jrnl.": []string{ "journal", },
		
		"Jrnls.": []string{ "journals", },
		
		"Jud.": []string{ "(in the Bible) Judges", },
		
		"Judg.": []string{ "(in the Bible) Judges", },
		
		"Jul.": []string{ "July", },
		
		"Jun.": []string{ "junior", },
		
		"Jurisd.": []string{ "jurisdiction", },
		
		"Jurisdict.": []string{ "jurisdiction", },
		
		"Jurispr.": []string{ "jurisprudence", },
		
		"Justif.": []string{ "justification", },
		
		"Justific.": []string{ "justification", },
		
		"K.": []string{ "king", },
		
		"Kent.": []string{ "Kentish", },
		
		"Kgs.": []string{ "(in the Bible) Kings", },
		
		"Kingd.": []string{ "kingdom", },
		
		"King’s Bench Div.": []string{ "King’s Bench Division", },
		
		"Knowl.": []string{ "knowledge", },
		
		"Kpr.": []string{ "keeper", },
		
		"LXX": []string{ "Septuagint", },
		
		"Lab.": []string{ "laboratory", },
		
		"Lam.": []string{ "(in the Bible) Lamentations", },
		
		"Lament": []string{ "lamentation", },
		
		"Lament.": []string{ "(in the Bible) Lamentations", },
		
		"Lanc.": []string{ "Lancashire", },
		
		"Lancash.": []string{ "Lancashire", },
		
		"Lancs.": []string{ "Lancashire", },
		
		"Lang.": []string{ "language", },
		
		"Langs.": []string{ "languages", },
		
		"Lat.": []string{ "Latin", },
		
		"Ld.": []string{ "Lord", },
		
		"Lds.": []string{ "Lords", },
		
		"Lect.": []string{ "lecture(s)", },
		
		"Leechd.": []string{ "leechdom", },
		
		"Leg.": []string{ "legend(s)", },
		
		"Leicest.": []string{ "Leicestershire", },
		
		"Leicester.": []string{ "Leicestershire", },
		
		"Leicestersh.": []string{ "Leicestershire", },
		
		"Leics.": []string{ "Leicestershire", },
		
		"Let.": []string{ "letter", },
		
		"Lett.": []string{ "letters", },
		
		"Lev.": []string{ "(in the Bible) Leviticus", },
		
		"Lex.": []string{ "lexicon", },
		
		"Libr.": []string{ "library", },
		
		"Limnol.": []string{ "limnology", },
		
		"Lincolnsh.": []string{ "Lincolnshire", },
		
		"Lincs.": []string{ "Lincolnshire", },
		
		"Ling.": []string{ "linguistic(s)", },
		
		"Linn.": []string{ "Linnean", },
		
		"Lit.": []string{ "literary"," literature", },
		
		"Lithogr.": []string{ "lithography", },
		
		"Lithol.": []string{ "lithology", },
		
		"Liturg.": []string{ "liturgy", },
		
		"Lond.": []string{ "London", },
		
		"MS.": []string{ "manuscript", },
		
		"MSS.": []string{ "manuscripts", },
		
		"Macc.": []string{ "(in the Apocrypha) Maccabees", },
		
		"Mach.": []string{ "machinery"," machine(s)", },
		
		"Mag.": []string{ "magazine", },
		
		"Magn.": []string{ "magnetism"," magnetic", },
		
		"Mal.": []string{ "(in the Bible) Malachi", },
		
		"Man.": []string{ "manual", },
		
		"Managem.": []string{ "management", },
		
		"Manch.": []string{ "Manchester", },
		
		"Manip.": []string{ "manipulation", },
		
		"Manuf.": []string{ "manufacturing"," manufacture(s)", },
		
		"Mar.": []string{ "March", },
		
		"Mass.": []string{ "Massachussets", },
		
		"Math.": []string{ "mathematics"," mathematical", },
		
		"Matt.": []string{ "(in the Bible) Matthew", },
		
		"Meas.": []string{ "measure", },
		
		"Measurem.": []string{ "measurement", },
		
		"Mech.": []string{ "mechanics"," mechanical", },
		
		"Med.": []string{ "medicine"," medical", },
		
		"Medit.": []string{ "meditation", },
		
		"Mem.": []string{ "memoir(s)", },
		
		"Merc.": []string{ "mercury", },
		
		"Merch.": []string{ "merchant(s)", },
		
		"Metall.": []string{ "metallurgy", },
		
		"Metallif.": []string{ "metalliferous", },
		
		"Metallogr.": []string{ "metallography", },
		
		"Metamorph.": []string{ "metamorphosis", },
		
		"Metaph.": []string{ "metaphysics", },
		
		"Meteorol.": []string{ "meteorology"," meteorologic(al)", },
		
		"Meth.": []string{ "method(s)", },
		
		"Metrop.": []string{ "metropolitan"," metropolis", },
		
		"Mex.": []string{ "Mexican", },
		
		"Mic.": []string{ "(in the Bible) Micah", },
		
		"Mich.": []string{ "Michigan", },
		
		"Microbiol.": []string{ "microbiology", },
		
		"Microsc.": []string{ "microscopic(al)", },
		
		"Mil.": []string{ "military", },
		
		"Milit.": []string{ "military", },
		
		"Min.": []string{ "mineralogy", },
		
		"Mineral.": []string{ "mineralogy"," minaralogic(al)", },
		
		"Misc.": []string{ "miscellaneous"," miscellany", },
		
		"Miscell.": []string{ "miscellany", },
		
		"Miss.": []string{ "Miss", },
		
		"Mod.": []string{ "modern", },
		
		"Monum.": []string{ "monument", },
		
		"Morphol.": []string{ "morphology", },
		
		"Mr.": []string{ "Mister", },
		
		"Mrs.": []string{ "Misses", },
		
		"Mt.": []string{ "mount(ain)", },
		
		"Mtg.": []string{ "meeting", },
		
		"Mts.": []string{ "mountains", },
		
		"Munic.": []string{ "municipal", },
		
		"Munif.": []string{ "munificence", },
		
		"Munim.": []string{ "muniment", },
		
		"Mus.": []string{ "museum", },
		
		"Myst.": []string{ "mystery", },
		
		"Myth.": []string{ "mythology", },
		
		"Mythol.": []string{ "mythology", },
		
		"N.": []string{ "north"," northern", },
		
		"N. Afr.": []string{ "North Africa(n)", },
		
		"N. Amer.": []string{ "North America(n)", },
		
		"N. Carolina": []string{ "North Carolina", },
		
		"N. Dakota": []string{ "North Dakota", },
		
		"N. Ir.": []string{ "Northern Irish", },
		
		"N. Irel.": []string{ "Northern Ireland", },
		
		"N.E.": []string{ "north-east"," north-eastern", },
		
		"N.E.D.": []string{ "New English Dictionary", },
		
		"N.S. Wales": []string{ "New South Wales", },
		
		"N.S.W.": []string{ "New South Wales", },
		
		"N.T.": []string{ "New Testament", },
		
		"N.W.": []string{ "north-west"," north-western", },
		
		"N.Y.": []string{ "New York", },
		
		"N.Z.": []string{ "New Zealand", },
		
		"Nah.": []string{ "(in the Bible) Nahum", },
		
		"Narr.": []string{ "narrative", },
		
		"Narrat.": []string{ "narrative", },
		
		"Nat.": []string{ "natural", },
		
		"Nat. Hist.": []string{ "natural history", },
		
		"Nat. Philos.": []string{ "natural philosophy", },
		
		"Nat. Sci.": []string{ "natural science", },
		
		"Naut.": []string{ "nautical", },
		
		"Nav.": []string{ "navigation", },
		
		"Navig.": []string{ "navigation", },
		
		"Neh.": []string{ "(in the Bible) Nehemiah", },
		
		"Neighb.": []string{ "neighbour(s)"," neighbourhood", },
		
		"Nerv.": []string{ "nervous", },
		
		"Neurol.": []string{ "neurology"," neurologic(al)", },
		
		"Neurosurg.": []string{ "neurosurgery", },
		
		"New Hampsh.": []string{ "New Hampshire", },
		
		"Newc.": []string{ "Newcastle", },
		
		"Newspr.": []string{ "newsprint", },
		
		"No.": []string{ "number", },
		
		"Non-conf.": []string{ "nonconformist(s)", },
		
		"Nonconf.": []string{ "nonconformist(s)", },
		
		"Norf.": []string{ "Norfolk", },
		
		"Northamptonsh.": []string{ "Northamptonshire", },
		
		"Northants.": []string{ "Northamptonshire", },
		
		"Northumb.": []string{ "Northumberland"," Northumbrian", },
		
		"Northumbld.": []string{ "Northumberland", },
		
		"Northumbr.": []string{ "Northumbrian", },
		
		"Norw.": []string{ "Norwegian", },
		
		"Norweg.": []string{ "Norwegian", },
		
		"Notts.": []string{ "Nottinghamshire", },
		
		"Nov.": []string{ "November", },
		
		"Nucl.": []string{ "nuclear", },
		
		"Num.": []string{ "(in the Bible) Numbers", },
		
		"Numism.": []string{ "numismatics", },
		
		"O.E.D.": []string{ "Oxford English Dictionary", },
		
		"O.T.": []string{ "Old Testament", },
		
		"OE": []string{ "(in dates) Old English", },
		
		"Obad.": []string{ "(in the Bible) Obadiah", },
		
		"Obed.": []string{ "obedience", },
		
		"Obj.": []string{ "object(s)", },
		
		"Obs.": []string{ "obsolete", },
		
		"Observ.": []string{ "observation(s)", },
		
		"Obstet.": []string{ "obstetric(s)", },
		
		"Obstetr.": []string{ "obstetric(s)", },
		
		"Obstetr. Med.": []string{ "obstetric medicine", },
		
		"Occas.": []string{ "occasional", },
		
		"Occup.": []string{ "occupation(al)", },
		
		"Occurr.": []string{ "occurrence", },
		
		"Oceanogr.": []string{ "oceanography"," oceanographic(al)", },
		
		"Oct.": []string{ "October", },
		
		"Off.": []string{ "official"," office", },
		
		"Offic.": []string{ "official", },
		
		"Okla.": []string{ "Oklahoma", },
		
		"Ont.": []string{ "Ontario", },
		
		"Ophthalm.": []string{ "ophthalmology", },
		
		"Ophthalmol.": []string{ "ophthalmology"," ophtalmologic(al)", },
		
		"Oppress.": []string{ "oppression", },
		
		"Opt.": []string{ "optics", },
		
		"Orac.": []string{ "oracle(s)", },
		
		"Ord.": []string{ "ordinance(s)", },
		
		"Org.": []string{ "organic", },
		
		"Org. Chem.": []string{ "organic chemistry", },
		
		"Organ. Chem.": []string{ "organic chemistry", },
		
		"Orig.": []string{ "original", },
		
		"Orkn.": []string{ "Orkney", },
		
		"Ornith.": []string{ "ornithology", },
		
		"Ornithol.": []string{ "ornithology"," ornithologic(al)", },
		
		"Orthogr.": []string{ "orthography", },
		
		"Outl.": []string{ "outline", },
		
		"Oxf.": []string{ "Oxford", },
		
		"Oxfordsh.": []string{ "Oxfordshire", },
		
		"Oxon.": []string{ "Oxfordshire", },
		
		"P. R.": []string{ "public records", },
		
		"Pa.": []string{ "Pennsylvania", },
		
		"Palaeobot.": []string{ "palaeobotany", },
		
		"Palaeogr.": []string{ "palaeography", },
		
		"Palaeont.": []string{ "palaeontology", },
		
		"Palaeontol.": []string{ "palaeontology"," palaeontologic(al)", },
		
		"Paraphr.": []string{ "paraphrase"," paraphrased", },
		
		"Parasitol.": []string{ "parasitology", },
		
		"Parl.": []string{ "parliament"," parliamentary", },
		
		"Parnass.": []string{ "Parnassus", },
		
		"Path.": []string{ "pathology"," pathologic(al)", },
		
		"Pathol.": []string{ "pathology"," pathologic(al)", },
		
		"Peculat.": []string{ "peculation(s)", },
		
		"Penins.": []string{ "peninsula", },
		
		"Perf.": []string{ "perfect"," perfection", },
		
		"Periodontol.": []string{ "periodontology"," periodontologic(al)", },
		
		"Pers.": []string{ "personal", },
		
		"Persec.": []string{ "persecution(s)", },
		
		"Perthsh.": []string{ "Perthshire", },
		
		"Pet.": []string{ "(in the Bible) Peter", },
		
		"Petrogr.": []string{ "petrography"," petrographic(al)", },
		
		"Petrol.": []string{ "petrology"," petrologic(al)", },
		
		"Pharm.": []string{ "pharmacology", },
		
		"Pharmaceut.": []string{ "pharmaceutical", },
		
		"Pharmacol.": []string{ "pharmacology"," pharmacological", },
		
		"Phil.": []string{ "(in the Bible) Philippians", },
		
		"Philad.": []string{ "Philadelphia", },
		
		"Philem.": []string{ "(in the Bible) Philemon", },
		
		"Philipp.": []string{ "(in the Bible) Philippians", },
		
		"Philol.": []string{ "philology"," philologic(al)", },
		
		"Philos.": []string{ "philosophy"," philosophic(al)", },
		
		"Phoen.": []string{ "Phoenician", },
		
		"Phonol.": []string{ "phonology", },
		
		"Photog.": []string{ "photography", },
		
		"Photogr.": []string{ "photography"," photographic(al)", },
		
		"Phrenol.": []string{ "phrenology", },
		
		"Phys.": []string{ "physiology", },
		
		"Physical Chem.": []string{ "physical chemistry", },
		
		"Physical Geogr.": []string{ "physical geography", },
		
		"Physiogr.": []string{ "physiography", },
		
		"Physiol.": []string{ "physiology"," physiologic(al)", },
		
		"Pict.": []string{ "picture(s)", },
		
		"Poet.": []string{ "poetical", },
		
		"Pol.": []string{ "politics"," political", },
		
		"Pol. Econ.": []string{ "political economy", },
		
		"Polit.": []string{ "politics"," political", },
		
		"Polytechn.": []string{ "polytechnic(al)", },
		
		"Pop.": []string{ "popular", },
		
		"Porc.": []string{ "porcelain", },
		
		"Port.": []string{ "Portuguese", },
		
		"Posth.": []string{ "posthumous(ly)", },
		
		"Postm.": []string{ "postmaster", },
		
		"Pott.": []string{ "pottery", },
		
		"Pract.": []string{ "practice", },
		
		"Predict.": []string{ "prediction", },
		
		"Pref.": []string{ "preface", },
		
		"Preh.": []string{ "prehistoric", },
		
		"Prehist.": []string{ "prehistory"," prehistoric", },
		
		"Prerog.": []string{ "prerogative", },
		
		"Pres.": []string{ "present", },
		
		"Presb.": []string{ "presbytery", },
		
		"Preserv.": []string{ "preservation", },
		
		"Prim.": []string{ "primer(s)", },
		
		"Princ.": []string{ "principle(s)", },
		
		"Print.": []string{ "printing", },
		
		"Probab.": []string{ "probability"," probabilities", },
		
		"Probl.": []string{ "problem", },
		
		"Proc.": []string{ "proceedings", },
		
		"Prod.": []string{ "product(s)", },
		
		"Prol.": []string{ "prologue", },
		
		"Pronunc.": []string{ "pronunciation", },
		
		"Prop.": []string{ "property"," properties", },
		
		"Pros.": []string{ "prosody", },
		
		"Prov.": []string{ "(in the Bible) Proverbs", },
		
		"Provid.": []string{ "providence(s)", },
		
		"Provinc.": []string{ "provincial", },
		
		"Provis.": []string{ "provisional", },
		
		"Ps.": []string{ "(in the Bible) Psalms", },
		
		"Psych.": []string{ "psychology", },
		
		"Psychoanal.": []string{ "psychoanalysis", },
		
		"Psychoanalyt.": []string{ "psychoanalytic", },
		
		"Psychol.": []string{ "psychology"," psychologic(al)", },
		
		"Psychopathol.": []string{ "psychopathology"," psychopathologic(al)", },
		
		"Pt.": []string{ "part", },
		
		"Publ.": []string{ "publication", },
		
		"Purg.": []string{ "purgatory", },
		
		"Q.": []string{ "queen", },
		
		"Q. Eliz.": []string{ "Queen Elizabeth", },
		
		"Qld.": []string{ "Queensland", },
		
		"Quantum Mech.": []string{ "quantum mechanics", },
		
		"Queen’s Bench Div.": []string{ "Queen’s Bench Division", },
		
		"R.": []string{ "royal", },
		
		"R.A.F.": []string{ "Royal Air Force", },
		
		"R.C.": []string{ "Roman Catholic(ism)", },
		
		"R.C. Church": []string{ "Roman Catholic Church", },
		
		"R.N.": []string{ "Royal Navy", },
		
		"Radiol.": []string{ "radiology", },
		
		"Reas.": []string{ "reason(s)", },
		
		"Reb.": []string{ "rebellion", },
		
		"Rebell.": []string{ "rebellion", },
		
		"Rec.": []string{ "record(s)", },
		
		"Reclam.": []string{ "reclamation", },
		
		"Recoll.": []string{ "recollection", },
		
		"Redempt.": []string{ "redemption", },
		
		"Ref.": []string{ "reformation", },
		
		"Refl.": []string{ "reflection(s)", },
		
		"Refus.": []string{ "refusal", },
		
		"Refut.": []string{ "refutation", },
		
		"Reg.": []string{ "register", },
		
		"Regic.": []string{ "regicide", },
		
		"Regist.": []string{ "register", },
		
		"Regr.": []string{ "register", },
		
		"Rel.": []string{ "relating", },
		
		"Relig.": []string{ "religion"," religious", },
		
		"Reminisc.": []string{ "reminiscence(s)", },
		
		"Remonstr.": []string{ "remonstrance", },
		
		"Renfrewsh.": []string{ "Renfrewshire", },
		
		"Rep.": []string{ "report", },
		
		"Reprod.": []string{ "reproduction"," reproductive", },
		
		"Rept.": []string{ "report", },
		
		"Repub.": []string{ "republic", },
		
		"Res.": []string{ "research", },
		
		"Resid.": []string{ "residence"," residential", },
		
		"Ret.": []string{ "return", },
		
		"Retrosp.": []string{ "retrospective", },
		
		"Rev.": []string{ "(in the Bible) Revelation", },
		
		"Revol.": []string{ "revolution", },
		
		"Rhet.": []string{ "rhetoric", },
		
		"Rhode Isl.": []string{ "Rhode Island", },
		
		"Rich.": []string{ "Richard", },
		
		"Rom.": []string{ "(in the Bible) Romans", },
		
		"Rom. Antiq.": []string{ "Roman antiquities", },
		
		"Ross-sh.": []string{ "Ross-shire", },
		
		"Roxb.": []string{ "Roxburgh", },
		
		"Roy.": []string{ "royal", },
		
		"Rudim.": []string{ "rudiment(s)", },
		
		"Russ.": []string{ "Russia"," Russian", },
		
		"S.": []string{ "south"," southern", },
		
		"S. Afr.": []string{ "South Africa(n)", },
		
		"S. Carolina": []string{ "South Carolina", },
		
		"S. Dakota": []string{ "South Dakota", },
		
		"S.E.": []string{ "south-east"," south-eastern", },
		
		"S.T.S.": []string{ "Scottish Text Society", },
		
		"S.W.": []string{ "south-west"," south-western", },
		
		"SS.": []string{ "saints", },
		
		"Sam.": []string{ "(in the Bible) Samuel", },
		
		"Sask.": []string{ "Saskatchewan", },
		
		"Sat.": []string{ "Saturday", },
		
		"Sax.": []string{ "Saxon", },
		
		"Sc.": []string{ "Scottish"," Scots", },
		
		"Scand.": []string{ "Scandinavia"," Scandinavian", },
		
		"Sch.": []string{ "school", },
		
		"Sci.": []string{ "science(s)"," scientific", },
		
		"Scot.": []string{ "Scottish", },
		
		"Scotl.": []string{ "Scotland", },
		
		"Script.": []string{ "scripture", },
		
		"Sculpt.": []string{ "sculpture", },
		
		"Seismol.": []string{ "seismology", },
		
		"Sel.": []string{ "selected", },
		
		"Sel. comm.": []string{ "select committee", },
		
		"Select.": []string{ "selection", },
		
		"Sept.": []string{ "September", },
		
		"Ser.": []string{ "series", },
		
		"Serm.": []string{ "sermon(s)", },
		
		"Sess.": []string{ "session", },
		
		"Settlem.": []string{ "settlement", },
		
		"Sev.": []string{ "several", },
		
		"Shakes.": []string{ "Shakespeare", },
		
		"Shaks.": []string{ "Shakespeare", },
		
		"Sheph.": []string{ "shepherd", },
		
		"Shetl.": []string{ "Shetland", },
		
		"Shropsh.": []string{ "Shropshire", },
		
		"Soc.": []string{ "society", },
		
		"Sociol.": []string{ "sociology"," sociological", },
		
		"Som.": []string{ "Somerset", },
		
		"Song Sol.": []string{ "(in the Bible) Song of Solomon", },
		
		"Song of Sol.": []string{ "(in the Bible) Song of Solomon", },
		
		"Sonn.": []string{ "sonnet", },
		
		"Span.": []string{ "Spanish", },
		
		"Spec.": []string{ "specimen(s)", },
		
		"Specif.": []string{ "specification(s)", },
		
		"Specim.": []string{ "specimen(s)", },
		
		"Spectrosc.": []string{ "spectroscopy", },
		
		"St.": []string{ "street", },
		
		"Staff.": []string{ "Staffordshire", },
		
		"Stafford.": []string{ "Staffordshire", },
		
		"Staffordsh.": []string{ "Staffordshire", },
		
		"Staffs.": []string{ "Staffordshire", },
		
		"Stand.": []string{ "standard", },
		
		"Stat.": []string{ "statistical", },
		
		"Statist.": []string{ "statistical", },
		
		"Stock Exch.": []string{ "stock exchange", },
		
		"Stratigr.": []string{ "stratigraphy"," stratigraphical", },
		
		"Struct.": []string{ "structure"," structural", },
		
		"Stud.": []string{ "studies", },
		
		"Subj.": []string{ "subject(s)", },
		
		"Subscr.": []string{ "subscription", },
		
		"Subscript.": []string{ "subscription", },
		
		"Suff.": []string{ "Suffolk", },
		
		"Suppl.": []string{ "supplement"," supplementary", },
		
		"Supplic.": []string{ "supplication(s)", },
		
		"Suppress.": []string{ "suppression", },
		
		"Surg.": []string{ "surgery"," surgical", },
		
		"Surv.": []string{ "survey"," surveying", },
		
		"Sus.": []string{ "(in the Apocrypha) Susanna", },
		
		"Symmetr.": []string{ "symmetrical", },
		
		"Symp.": []string{ "symposium", },
		
		"Syst.": []string{ "system", },
		
		"Taxon.": []string{ "taxonomy"," taxonomic", },
		
		"Techn.": []string{ "technical", },
		
		"Technol.": []string{ "technology"," technological", },
		
		"Tel.": []string{ "telegraph", },
		
		"Telecomm.": []string{ "telecommunications", },
		
		"Telegr.": []string{ "telegraphy"," telegraphic(al)", },
		
		"Teleph.": []string{ "telephony"," telephonic", },
		
		"Teratol.": []string{ "teratology", },
		
		"Terminol.": []string{ "terminology"," terminological", },
		
		"Terrestr.": []string{ "terrestrial", },
		
		"Test.": []string{ "testament", },
		
		"Textbk.": []string{ "textbook", },
		
		"Theat.": []string{ "theatre"," theatrical", },
		
		"Theatr.": []string{ "theatre"," theatrical", },
		
		"Theol.": []string{ "theology"," theological", },
		
		"Theoret.": []string{ "theoretical", },
		
		"Thermonucl.": []string{ "thermonuclear", },
		
		"Thes.": []string{ "thesaurus", },
		
		"Thess.": []string{ "(in the Bible) Thessalonians", },
		
		"Tim.": []string{ "(in the Bible) Timothy", },
		
		"Tit.": []string{ "(in the Bible) Titus", },
		
		"Topogr.": []string{ "topography"," topographic(al)", },
		
		"Trad.": []string{ "tradition"," traditional", },
		
		"Trag.": []string{ "tragedy", },
		
		"Trans.": []string{ "transaction(s)", },
		
		"Transl.": []string{ "translation", },
		
		"Transubstant.": []string{ "transubstantiation", },
		
		"Trav.": []string{ "travel(s)", },
		
		"Treas.": []string{ "treasurer"," treasury", },
		
		"Treat.": []string{ "treatise", },
		
		"Treatm.": []string{ "treatment", },
		
		"Trib.": []string{ "tribune", },
		
		"Trig.": []string{ "trigonometry", },
		
		"Trigonom.": []string{ "trigonometry", },
		
		"Trop.": []string{ "tropical", },
		
		"Troub.": []string{ "troublesome", },
		
		"Troubl.": []string{ "troublesome", },
		
		"Typog.": []string{ "typography", },
		
		"Typogr.": []string{ "typography", },
		
		"U.K.": []string{ "United Kingdom", },
		
		"U.S.": []string{ "United States", },
		
		"U.S.A.F.": []string{ "United States Air Force", },
		
		"U.S.S.R.": []string{ "Union of Soviet Socialist Republics", },
		
		"Univ.": []string{ "university", },
		
		"Unnat.": []string{ "unnatural", },
		
		"Unoffic.": []string{ "unofficial", },
		
		"Urin.": []string{ "urinary", },
		
		"Utilit.": []string{ "utilitarian", },
		
		"Va.": []string{ "Virginia", },
		
		"Vac.": []string{ "vacation", },
		
		"Valedict.": []string{ "valediction", },
		
		"Veg.": []string{ "vegetable", },
		
		"Veg. Phys.": []string{ "vegetable physiology", },
		
		"Veg. Physiol.": []string{ "vegetable physiology", },
		
		"Venet.": []string{ "Venetian", },
		
		"Vertebr.": []string{ "vertebrate", },
		
		"Vet.": []string{ "veterinary"," veterinary science", },
		
		"Vet. Med.": []string{ "veterinary medicine", },
		
		"Vet. Path.": []string{ "veterinary pathology", },
		
		"Vet. Sci.": []string{ "veterinary science", },
		
		"Vet. Surg.": []string{ "veterinary surgery", },
		
		"Vic.": []string{ "Victoria", },
		
		"Vict.": []string{ "Victoria", },
		
		"Vind.": []string{ "vindication"," vindicated", },
		
		"Vindic.": []string{ "vindication"," vindicated", },
		
		"Virg.": []string{ "virgin(s)", },
		
		"Virol.": []string{ "virology", },
		
		"Voc.": []string{ "vocabulary", },
		
		"Vocab.": []string{ "vocabulary", },
		
		"Vol.": []string{ "volume", },
		
		"Vols.": []string{ "volumes", },
		
		"Voy.": []string{ "voyage(s)", },
		
		"Vulg.": []string{ "Vulgate", },
		
		"W.": []string{ "west"," western", },
		
		"W. Afr.": []string{ "West Africa(n)", },
		
		"W. Ind.": []string{ "West Indies"," West Indian", },
		
		"W. Indies": []string{ "West Indies", },
		
		"W. Va.": []string{ "West Virginia", },
		
		"Warwicksh.": []string{ "Warwickshire", },
		
		"Wd.": []string{ "word", },
		
		"Westm.": []string{ "Westminster", },
		
		"Westmld.": []string{ "Westmoreland", },
		
		"Westmorld.": []string{ "Westmoreland", },
		
		"Westmrld.": []string{ "Westmoreland", },
		
		"Will.": []string{ "William", },
		
		"Wilts.": []string{ "Wiltshire", },
		
		"Wiltsh.": []string{ "Wiltshire", },
		
		"Wis.": []string{ "Wisconsin", },
		
		"Wisd.": []string{ "(in the Apocrypha) Wisdom of Solomon", },
		
		"Wk.": []string{ "work", },
		
		"Wkly.": []string{ "weekly", },
		
		"Wks.": []string{ "works", },
		
		"Wonderf.": []string{ "wonderful", },
		
		"Worc.": []string{ "Worcestershire", },
		
		"Worcestersh.": []string{ "Worcestershire", },
		
		"Worcs.": []string{ "Worcestershire", },
		
		"Writ.": []string{ "writing(s)", },
		
		"Yearbk.": []string{ "yearbook", },
		
		"Yng.": []string{ "young", },
		
		"Yorks.": []string{ "Yorkshire", },
		
		"Yorksh.": []string{ "Yorkshire", },
		
		"Yr.": []string{ "year", },
		
		"Yrs.": []string{ "years", },
		
		"Zech.": []string{ "(in the Bible) Zechariah", },
		
		"Zeitschr.": []string{ "Zeitschrift", },
		
		"Zeph.": []string{ "(in the Bible) Zephaniah", },
		
		"Zoogeogr.": []string{ "zoogeography", },
		
		"Zool.": []string{ "zoology"," zoological", },
		
		"a": []string{ "(in dates) ante", },
		
		"abbrev.": []string{ "abbreviation (of)", },
		
		"abl.": []string{ "ablative", },
		
		"absol.": []string{ "absolute(ly)", },
		
		"abstr.": []string{ "abstract", },
		
		"acc.": []string{ "accusative", },
		
		"accus.": []string{ "accusative", },
		
		"ad.": []string{ "adaptation of", },
		
		"adj.": []string{ "adjective"," adjectival", },
		
		"adj. phr.": []string{ "adjectival phrase", },
		
		"adjs.": []string{ "adjectives", },
		
		"adv.": []string{ "adverb"," adverbial(ly)", },
		
		"advb.": []string{ "adverb"," adverbial(ly)", },
		
		"advs.": []string{ "adverbs", },
		
		"agst.": []string{ "against", },
		
		"alt.": []string{ "alteration", },
		
		"aphet.": []string{ "aphetic"," aphetized", },
		
		"app.": []string{ "apparently", },
		
		"appos.": []string{ "appositive(ly)", },
		
		"arch.": []string{ "archaic", },
		
		"art.": []string{ "article", },
		
		"attrib.": []string{ "attributive(ly)", },
		
		"bef.": []string{ "before", },
		
		"betw.": []string{ "between", },
		
		"c": []string{ "(in dates) circa", },
		
		"c.": []string{ "century", },
		
		"cent.": []string{ "century", },
		
		"cf.": []string{ "confer"," ‘compare’", },
		
		"cl.": []string{ "clause", },
		
		"cogn. w.": []string{ "cognate with", },
		
		"collect.": []string{ "collective(ly)", },
		
		"colloq.": []string{ "colloquial(ly)", },
		
		"comb. form": []string{ "combining form", },
		
		"comp.": []string{ "compound"," composition", },
		
		"compar.": []string{ "comparative", },
		
		"compl.": []string{ "complement", },
		
		"conc.": []string{ "concerning", },
		
		"concr.": []string{ "concrete(ly)", },
		
		"conj.": []string{ "conjunction", },
		
		"cons.": []string{ "consonant", },
		
		"const.": []string{ "construction"," construed with", },
		
		"contempt.": []string{ "contemptuous(ly)", },
		
		"contr.": []string{ "contrasted (with)", },
		
		"corresp.": []string{ "corresponding (to)", },
		
		"cpd.": []string{ "compound", },
		
		"dat.": []string{ "dative", },
		
		"def.": []string{ "definition", },
		
		"dem.": []string{ "demonstrative", },
		
		"deriv.": []string{ "derivative (in etymologies)", },
		
		"derog.": []string{ "derogatory", },
		
		"dial.": []string{ "dialect", },
		
		"dim.": []string{ "diminutive", },
		
		"dyslog.": []string{ "dyslogistic", },
		
		"e.": []string{ "east", },
		
		"e. midl.": []string{ "east midland (dialect)", },
		
		"eOE": []string{ "(in dates) early Old English", },
		
		"east.": []string{ "eastern (dialect)", },
		
		"ed.": []string{ "edition", },
		
		"ellipt.": []string{ "elliptical(ly)", },
		
		"emph.": []string{ "emphatic(ally)", },
		
		"erron.": []string{ "erroneous(ly)", },
		
		"esp.": []string{ "especially", },
		
		"etym.": []string{ "etymology", },
		
		"etymol.": []string{ "etymological(ly)", },
		
		"euphem.": []string{ "euphemistic(ally)", },
		
		"exc.": []string{ "except", },
		
		"f.": []string{ "folio", },
		
		"fam.": []string{ "familiar(ly)", },
		
		"famil.": []string{ "familiar(ly)", },
		
		"fem.": []string{ "feminine", },
		
		"fig.": []string{ "figurative(ly)", },
		
		"fl.": []string{ "floruit"," ‘flourished’", },
		
		"freq.": []string{ "frequently", },
		
		"fut.": []string{ "future", },
		
		"gen.": []string{ "genitive", },
		
		"gerund.": []string{ "gerundial", },
		
		"hist.": []string{ "historical", },
		
		"imit.": []string{ "imitative", },
		
		"imp.": []string{ "imperative", },
		
		"imperf.": []string{ "imperfect", },
		
		"impers.": []string{ "impersonal", },
		
		"impf.": []string{ "imperfect", },
		
		"improp.": []string{ "improper(ly)", },
		
		"ind.": []string{ "indirect", },
		
		"indef.": []string{ "indefinite", },
		
		"indic.": []string{ "indicative", },
		
		"indir.": []string{ "indirect", },
		
		"infin.": []string{ "infinitive", },
		
		"infl.": []string{ "influenced", },
		
		"instr.": []string{ "instrumental", },
		
		"int.": []string{ "interjection", },
		
		"interj.": []string{ "interjection", },
		
		"interrog.": []string{ "interrogative", },
		
		"intr.": []string{ "intransitive", },
		
		"intrans.": []string{ "intransitive", },
		
		"iron.": []string{ "ironically", },
		
		"irreg.": []string{ "irregular(ly)", },
		
		"joc.": []string{ "jocular(ly)", },
		
		"l.": []string{ "line", },
		
		"lOE": []string{ "(in dates) late Old English", },
		
		"lit.": []string{ "literal(ly)", },
		
		"ll.": []string{ "lines", },
		
		"m.": []string{ "masculine", },
		
		"masc.": []string{ "masculine", },
		
		"med.": []string{ "medieval", },
		
		"metaphor.": []string{ "metaphorical(ly)", },
		
		"metr. gr.": []string{ "metri gratia", },
		
		"midl.": []string{ "midland(s) (dialect)", },
		
		"mispr.": []string{ "misprinted", },
		
		"mod.": []string{ "modern", },
		
		"n.": []string{ "noun", },
		
		"n.e.": []string{ "north-eastern (dialect)", },
		
		"n.w.": []string{ "north-western (dialect)", },
		
		"nom.": []string{ "nominative", },
		
		"nonce-wd.": []string{ "nonce-word", },
		
		"north.": []string{ "northern (dialect)", },
		
		"ns.": []string{ "nouns", },
		
		"obj.": []string{ "object"," objective", },
		
		"obl.": []string{ "oblique", },
		
		"obs.": []string{ "obsolete", },
		
		"occas.": []string{ "occasionally", },
		
		"opp.": []string{ "opposed (to)"," the opposite (of)", },
		
		"orig.": []string{ "origin(al)(ly)", },
		
		"p.": []string{ "page", },
		
		"pa.": []string{ "past", },
		
		"pa. pple.": []string{ "passive participle"," past participle", },
		
		"pa. t.": []string{ "past tense", },
		
		"pass.": []string{ "passive(ly)", },
		
		"perf.": []string{ "perfect", },
		
		"perh.": []string{ "perhaps", },
		
		"pers.": []string{ "person(al)", },
		
		"personif.": []string{ "personified", },
		
		"pf.": []string{ "perfect", },
		
		"phonet.": []string{ "phonetic(ally)", },
		
		"phr.": []string{ "phrase", },
		
		"pl.": []string{ "plural", },
		
		"plur.": []string{ "plural", },
		
		"poet.": []string{ "poetic(al)", },
		
		"pop.": []string{ "popular(ly)", },
		
		"poss.": []string{ "possible"," possibly", },
		
		"ppl.": []string{ "participial", },
		
		"ppl. a.": []string{ "participial adjective", },
		
		"ppl. adj.": []string{ "participial adjective", },
		
		"ppl. adjs.": []string{ "participial adjectives", },
		
		"pple.": []string{ "participle", },
		
		"pples.": []string{ "participles", },
		
		"pr.": []string{ "present", },
		
		"pr. pple.": []string{ "present participle", },
		
		"prec.": []string{ "preceding", },
		
		"pred.": []string{ "predicate", },
		
		"predic.": []string{ "predicate"," predicative(ly)", },
		
		"pref.": []string{ "prefix", },
		
		"prep.": []string{ "preposition(al)", },
		
		"pres.": []string{ "present", },
		
		"pres. pple.": []string{ "present participle", },
		
		"priv.": []string{ "privative", },
		
		"prob.": []string{ "probably", },
		
		"pron.": []string{ "pronunciation", },
		
		"pronunc.": []string{ "pronunciation", },
		
		"prop.": []string{ "properly", },
		
		"propr.": []string{ "proprietary", },
		
		"prov.": []string{ "proverb(ial)(ly)", },
		
		"pseudo-Sc.": []string{ "pseudo-Scots", },
		
		"pseudo-arch.": []string{ "pseudo-archaic", },
		
		"pseudo-dial.": []string{ "pseudo-dialect", },
		
		"q.v.": []string{ "quod vide"," ‘which see’", },
		
		"quot.": []string{ "quotation", },
		
		"quots.": []string{ "quotations", },
		
		"redupl.": []string{ "reduplicating", },
		
		"refash.": []string{ "refashioned"," refashioning", },
		
		"refl.": []string{ "reflexive", },
		
		"reg.": []string{ "regular", },
		
		"rel.": []string{ "relative", },
		
		"repr.": []string{ "representative"," representing", },
		
		"rhet.": []string{ "rhetorical(ly)", },
		
		"s.": []string{ "south", },
		
		"s.e.": []string{ "south-eastern (dialect)", },
		
		"s.v.": []string{ "sub verbo"," ‘under the word’", },
		
		"s.w.": []string{ "south-western (dialect)", },
		
		"sc.": []string{ "scilicet"," ‘understand’ or ‘supply’", },
		
		"sing.": []string{ "singular", },
		
		"south.": []string{ "southern (dialect)", },
		
		"sp.": []string{ "spelling", },
		
		"spec.": []string{ "specifically", },
		
		"str.": []string{ "strong", },
		
		"subj.": []string{ "subjunctive", },
		
		"subjunct.": []string{ "subjunctive", },
		
		"subord.": []string{ "subordinate", },
		
		"subord. cl.": []string{ "subordinate clause", },
		
		"subseq.": []string{ "subsequent(ly)", },
		
		"subst.": []string{ "substantive(ly)", },
		
		"suff.": []string{ "suffix", },
		
		"superl.": []string{ "superlative", },
		
		"syll.": []string{ "syllable", },
		
		"t.": []string{ "tense", },
		
		"techn.": []string{ "technical", },
		
		"tr.": []string{ "translation"," translating", },
		
		"trans.": []string{ "transitive", },
		
		"transf.": []string{ "transferred sense", },
		
		"transl.": []string{ "translation"," translating", },
		
		"ult.": []string{ "ultimately", },
		
		"unkn.": []string{ "unknown", },
		
		"unstr.": []string{ "unstressed", },
		
		"usu.": []string{ "usually", },
		
		"v.": []string{ "verb", },
		
		"v.r.": []string{ "variant reading", },
		
		"v.rr.": []string{ "variant readings", },
		
		"var.": []string{ "variant", },
		
		"varr.": []string{ "variants", },
		
		"vars.": []string{ "variants", },
		
		"vb.": []string{ "verb", },
		
		"vbl.": []string{ "verbal", },
		
		"vbl. ns.": []string{ "verbal nouns", },
		
		"vbl.n.": []string{ "verbal noun", },
		
		"vbs.": []string{ "verbs", },
		
		"viz.": []string{ "videlicet"," ‘namely’", },
		
		"vulg.": []string{ "vulgar(ly)", },
		
		"w.": []string{ "west", },
		
		"wd.": []string{ "word", },
		
		"west.": []string{ "western", },
		
		"wk.": []string{ "weak", },
		
	}

	suffixes = map[string][]string{
		
		"Assoc.": []string{ "Football", },
		
		"Bel": []string{ "& Dr.", },
		
		"Ch.": []string{ "Hist.", },
		
		"Chem.": []string{ "Engin.", },
		
		"Civ.": []string{ "Law", },
		
		"Civil": []string{ "Engin.", },
		
		"Class.": []string{ "Antiq.", },
		
		"Comm.": []string{ "Law", },
		
		"Comp.": []string{ "Anat.", },
		
		"Constit.": []string{ "Hist.", },
		
		"Crim.": []string{ "Law", },
		
		"E.": []string{ "Afr.","Angl.","Anglian","Ind.", },
		
		"East": []string{ "Ind.", },
		
		"Eccl.": []string{ "Hist.","Law", },
		
		"Electr.": []string{ "Engin.", },
		
		"Encycl.": []string{ "Brit.","Metrop.", },
		
		"Ex.": []string{ "doc.", },
		
		"Gramm.": []string{ "Analysis", },
		
		"Industr.": []string{ "Rel.", },
		
		"Invert.": []string{ "Zool.", },
		
		"King’s": []string{ "Bench Div.", },
		
		"N.": []string{ "Afr.","Amer.","Carolina","Dakota","Ir.","Irel.", },
		
		"N.S.": []string{ "Wales", },
		
		"Nat.": []string{ "Hist.","Philos.","Sci.", },
		
		"New": []string{ "Hampsh.", },
		
		"Obstetr.": []string{ "Med.", },
		
		"Org.": []string{ "Chem.", },
		
		"Organ.": []string{ "Chem.", },
		
		"P.": []string{ "R.", },
		
		"Physical": []string{ "Chem.","Geogr.", },
		
		"Pol.": []string{ "Econ.", },
		
		"Q.": []string{ "Eliz.", },
		
		"Quantum": []string{ "Mech.", },
		
		"Queen’s": []string{ "Bench Div.", },
		
		"R.C.": []string{ "Church", },
		
		"Rhode": []string{ "Isl.", },
		
		"Rom.": []string{ "Antiq.", },
		
		"S.": []string{ "Afr.","Carolina","Dakota", },
		
		"Sel.": []string{ "comm.", },
		
		"Song": []string{ "of Sol.","Sol.", },
		
		"Stock": []string{ "Exch.", },
		
		"Veg.": []string{ "Phys.","Physiol.", },
		
		"Vet.": []string{ "Med.","Path.","Sci.","Surg.", },
		
		"W.": []string{ "Afr.","Ind.","Indies","Va.", },
		
		"adj.": []string{ "phr.", },
		
		"cogn.": []string{ "w.", },
		
		"comb.": []string{ "form", },
		
		"e.": []string{ "midl.", },
		
		"metr.": []string{ "gr.", },
		
		"pa.": []string{ "pple.","t.", },
		
		"ppl.": []string{ "a.","adj.","adjs.", },
		
		"pr.": []string{ "pple.", },
		
		"pres.": []string{ "pple.", },
		
		"subord.": []string{ "cl.", },
		
		"vbl.": []string{ "ns.", },
		
	}
)

// IsStrictlyAbbreviated returns true/false if giving text is an abbreviation.
// This function does not provide optimized truthy value based on prefix. It
// strictly validates if this is in itself a abbreviation.
func IsStrictlyAbbreviated(item string) bool {
	_, ok := abbreviations[item]
	return ok
}

// IsAbbreviated returns true/false if giving text is an abbreviation.
// IsAbbreviated provides a optimistic positive return value of true
// if giving abbreviation has other suffix like Song of Sol. etc.
func IsAbbreviated(item string) bool {
	// if we have direct hit, then return true.
	if _, ok := abbreviations[item]; ok {
		return ok
	}


	// We need to attempt indirect hit with combo map.
	// and return optimistic truthy value if any is found.
	if _, ok := suffixes[item]; ok {
		return ok
	}

	
	return false
}

// GetAllSuffix returns a slice of strings that represent
// suffixed types of giving abbreviation word if any exists.
// That is if you provide 'Song.', then we return ['Sol.', 'of Sol.'].
func GetAllSuffix(item string) ([]string, bool) {
	suffixes, ok := suffixes[item]
	return suffixes, ok
}

// Mutations returns a slice of strings that represent
// all possible mutation of giving abbreviation prefix if any.
// That is if you provide 'Song.', then we return ['Song. Sol.', 'Song of Sol.'].
func Mutations(item string) ([]string, bool) {
	if suffixes, ok := suffixes[item]; ok {
		var parts []string
		for _, suffix := range suffixes {
			parts = append(parts, item+" "+suffix)
		}
		return parts, true
	}

	return nil, false
}

// GetMeanings returns all available meanings of given abbreviation if found.
// Bool argument is used to indicate if abbreviation was found.
func GetMeanings(abbr string) ([]string, bool) {
	meanings, ok := abbreviations[abbr]
	return meanings, ok
}

