from translate import Translator
translator= Translator(from_lang="uzbek", to_lang="english")
translation=translator.translate("men olma yemayman")
print(translation)